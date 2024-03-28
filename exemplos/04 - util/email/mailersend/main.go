/*
	Criar uma conta em mailersend (conta free)
	pegar a chave da API e colocar na variavel APIKey

	Substitua os dados da mensagem, e pronto!

	Observação: para fins de teste, usar info@trial-k68zxl2e2mmlj905.mlsender.net no campo from

	Para maiores informações, consulte a documentação em https://developers.mailersend.com/

*/

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "api-key"

func main() {
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "Greetings from the team, you got this message through MailerSend."
	html := "Greetings from the team, you got this message through MailerSend."

	from := mailersend.From{
		Name:  "Your name or company",
		Email: "info@trial-k68zxl2e2mmlj905.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Recipient",
			Email: "Recipient@email.com",
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)

	res, err := ms.Email.Send(ctx, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(res.Header.Get("X-Message-Id"))

}
