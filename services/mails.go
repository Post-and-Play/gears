package services

import (
	"fmt"
	"net/smtp"
	"os"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/templates"
)

// SendMail
func SendMail(receiver *models.Receiver, mail *models.MailRequest) {
	var sender models.Sender

	sender.SenderMail = os.Getenv("SENDER_MAIL")
	sender.SenderPass = os.Getenv("SENDER_PASS")

	to := []string{
		receiver.ReceiverMail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", sender.SenderMail, sender.SenderPass, smtpHost)

	//fmt.Println(sender.SenderMail + "\n" + sender.SenderPass)

	t, body := templates.BuildTemplate(mail.Subject)

	t.Execute(&body, struct {
		Title string
		Message string
		Link string
		Footer string
		ButtonText string
	}{
		Title: mail.Title,
		Message: mail.Message,
		Link: mail.Link,
		Footer: mail.Footer,
		ButtonText: mail.ButtonText,
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender.SenderMail, to, body.Bytes())
	if err != nil {
		mail.OK = false
		fmt.Println(err)
		return 
	}

	mail.OK = true
	fmt.Println("isvalid: Email Sent!")
	return 
}
