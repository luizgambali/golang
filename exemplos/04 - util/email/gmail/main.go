// reference: https://www.courier.com/guides/golang-send-email/
package main

import (
	"net/smtp"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	SendMail()
}

// as rotinas abaixo não vão funcionar se o two-factor authentication estiver ativado
// para desativar acesse: https://myaccount.google.com/security-checkup

func SendMail() {
	from := "<paste your gmail account here>"
	password := "<paste Google password or app password here>"

	toEmailAddress := "<paste the email address you want to send to>"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}

func SendMail2() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "<paste your gmail account here>")
	msg.SetHeader("To", "<paste the email address you want to send to>")
	msg.SetHeader("Subject", "<paste the subject of the mail>")
	msg.SetBody("text/html", "<b>This is the body of the mail</b>")
	msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.gmail.com", 587, "<paste your gmail account here>", "<paste Google password or app password here>")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
