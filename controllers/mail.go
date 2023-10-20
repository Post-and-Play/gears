package controllers

import (
	"fmt"
	"net/smtp"
	"os"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/templates"
	"github.com/gin-gonic/gin"
)

// SendMail
func SendMail(c *gin.Context) {
	var sender models.Sender
	var receiver models.Receiver

	sender.SenderMail = os.Getenv("SENDER_MAIL")
	sender.SenderPass = os.Getenv("SENDER_PASS")

	to := []string{
		receiver.ReceiverMail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", sender.SenderMail, sender.SenderPass, smtpHost)

	t, body := templates.BuildTemplate("Message")

	t.Execute(&body, struct {
		Subject string
		Message string
	}{
		Subject: "Subject Email",
		Message: "Body Email",
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender.SenderMail, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
