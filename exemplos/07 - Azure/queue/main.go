// reference: https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue

package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

// Please set the ACCOUNT_NAME and ACCOUNT_KEY environment variables to your storage account's
// name and account key, before running the examples.
func accountInfo() (string, string) {
	return os.Getenv("ACCOUNT_NAME"), os.Getenv("ACCOUNT_KEY")
}

func main() {
	// From the Azure portal, get your Storage account's name and account key.
	accountName, accountKey := accountInfo()

	// Use your Storage account's name and key to create a credential object; this is used to access your account.
	credential, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create a request pipeline that is used to process HTTP(S) requests and responses. It requires
	// your account credentials. In more advanced scenarios, you can configure telemetry, retry policies,
	// logging, and other options. Also, you can configure multiple request pipelines for different scenarios.
	p := azqueue.NewPipeline(credential, azqueue.PipelineOptions{})

	// From the Azure portal, get your Storage account queue service URL endpoint.
	// The URL typically looks like this:
	u, _ := url.Parse(fmt.Sprintf("https://%s.queue.core.windows.net", accountName))

	// Create an ServiceURL object that wraps the service URL and a request pipeline.
	serviceURL := azqueue.NewServiceURL(*u, p)

	// Now, you can use the serviceURL to perform various queue operations.

	// All HTTP operations allow you to specify a Go context.Context object to control cancellation/timeout.
	ctx := context.TODO() // This example uses a never-expiring context.

	// Create a URL that references a queue in your Azure Storage account.
	// This returns a QueueURL object that wraps the queue's URL and a request pipeline (inherited from serviceURL)
	queueURL := serviceURL.NewQueueURL("examplequeue") // Queue names require lowercase

	// The code below shows how to create the queue. It is common to create a queue and never delete it:
	_, err = queueURL.Create(ctx, azqueue.Metadata{})
	if err != nil {
		log.Fatal(err)
	}

	// The code below shows how a client application enqueues 2 messages into the queue:
	// Create a URL allowing you to manipulate a queue's messages.
	// This returns a MessagesURL object that wraps the queue's messages URL and a request pipeline (inherited from queueURL)
	messagesURL := queueURL.NewMessagesURL()

	// Enqueue 2 messages
	_, err = messagesURL.Enqueue(ctx, "This is message 1", time.Second*0, time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	_, err = messagesURL.Enqueue(ctx, "This is message 2", time.Second*0, time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	// The code below shows how a client or server can determine the approximate count of messages in the queue:
	props, err := queueURL.GetProperties(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Approximate number of messages in the queue=%d\n", props.ApproximateMessagesCount())

	// The code below shows how to initialize a service that wishes to process messages:
	const concurrentMsgProcessing = 16 // Set this as you desire
	msgCh := make(chan *azqueue.DequeuedMessage, concurrentMsgProcessing)
	const poisonMessageDequeueThreshold = 4 // Indicates how many times a message is attempted to be processed before considering it a poison message

	// Create goroutines that can process messages in parallel
	for n := 0; n < concurrentMsgProcessing; n++ {
		go func(msgCh <-chan *azqueue.DequeuedMessage) {
			for {
				msg := <-msgCh // Get a message from the channel

				// Create a URL allowing you to manipulate this message.
				// This returns a MessageIDURL object that wraps the this message's URL and a request pipeline (inherited from messagesURL)
				msgIDURL := messagesURL.NewMessageIDURL(msg.ID)
				popReceipt := msg.PopReceipt // This message's most-recent pop receipt

				if msg.DequeueCount > poisonMessageDequeueThreshold {
					// This message has attempted to be processed too many times; treat it as a poison message
					// DO NOT attempt to process this message
					// Log this message as a poison message somewhere (code not shown)
					// Delete this poison message from the queue so it will never be dequeued again
					msgIDURL.Delete(ctx, popReceipt)
					continue // Process a different message
				}

				// This message is not a poison message, process it (this example just displays it):
				fmt.Print(msg.Text + "\n")

				// NOTE: You can examine/use any of the message's other properties as you desire:
				_, _, _ = msg.InsertionTime, msg.ExpirationTime, msg.NextVisibleTime

				// OPTIONAL: while processing a message, you can update the message's visibility timeout
				// (to prevent other servers from dequeuing the same message simultaneously) and update the
				// message's text (to prevent some successfully-completed processing from re-executing the
				// next time this message is dequeued):
				update, err := msgIDURL.Update(ctx, popReceipt, time.Second*20, "updated msg")
				if err != nil {
					log.Fatal(err)
				}
				popReceipt = update.PopReceipt // Performing any operation on a message ID always requires the most recent pop receipt

				// After processing the message, delete it from the queue so it won't be dequeued ever again:
				_, err = msgIDURL.Delete(ctx, popReceipt)
				if err != nil {
					log.Fatal(err)
				}
				// Loop around to process the next message
			}
		}(msgCh)
	}

	// The code below shows the service's infinite loop that dequeues messages and dispatches them in batches for processsing:
	for {
		// Try to dequeue a batch of messages from the queue
		dequeue, err := messagesURL.Dequeue(ctx, azqueue.QueueMaxMessagesDequeue, 10*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		if dequeue.NumMessages() == 0 {
			// The queue was empty; sleep a bit and try again
			// Shorter time means higher costs & less latency to dequeue a message
			// Higher time means lower costs & more latency to dequeue a message
			time.Sleep(time.Second * 10)
		} else {
			// We got some messages, put them in the channel so that many can be processed in parallel:
			// NOTE: The queue does not guarantee FIFO ordering & processing messages in parallel also does
			// not preserve FIFO ordering. So, the "Output:" order below is not guaranteed but usually works.
			for m := int32(0); m < dequeue.NumMessages(); m++ {
				msgCh <- dequeue.Message(m)
			}
		}
		// This batch of dequeued messages are in the channel, dequeue another batch
		break // NOTE: For this example only, break out of the infinite loop so this example terminates
	}

	time.Sleep(time.Second * 10) // For this example, delay in hopes that both messages complete processing before the example terminates

	// This example deletes the queue (to clean up) but normally, you never delete a queue.
	_, err = queueURL.Delete(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
