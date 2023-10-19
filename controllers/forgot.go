package controllers

import (
	"log"
	"net/http"
	"time"
	"github.com/Post-and-Play/gears/services"
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

// ForgotUser godoc
// @Summary      ForgotUser password
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.Forgot  true  "Forgot Model"
// @Success      200  {object}  models.Forgot
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /forgot [get]
func ForgotUser(c *gin.Context) {
	var user models.User
	email := c.Query("mail")

	infra.DB.Where("mail = ?", email).First(&user)

	if user.ID == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	currentTime := time.Now()

	user.SecurityKey = services.SHA256Encoder(user.Mail + user.Password + currentTime.String())

	if infra.DB.Model(&user).Updates(&user).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	/*
	client := infra.NewHttpClient(os.Getenv("MAIL_URL"))
	mailResponse, err := client.MailPost(mail)
	if err != nil {
		log.Default().Printf("Mail error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error": err.Error()})
		return
	}
	*/

	var receiver models.Receiver
	receiver.ReceiverMail = email 

	var mailRequest models.MailRequest
	mailRequest.Subject = "PAP Redefinicao de senha"
	mailRequest.Body =  "Ola Para redefinir sua senha clique no link abaixo"

	mailResponse, err := mails.SendMail(receiver, mailRequest)
	if err != nil {
		log.Default().Printf("Mail error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": mailResponse})
}




// RecoverPassword godoc
// @Summary      Recover password
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.RecoverPassword  true  "RecoverPassword Model"
// @Success      200  {object}  models.Edwiges
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /recover [post]
func RecoverPassword(c *gin.Context) {
	var recover models.RecoverPassword
	
	if err := c.ShouldBindJSON(&recover); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.RecoverPasswordValidator(&recover); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	/*client := infra.NewHttpClient(os.Getenv("EDWIGES_URL"))
	edwigesResponse, err := client.EdwigesPost(edwigesRequest.Mail)
	if err != nil {
		log.Default().Printf("Mail error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error": err.Error()})
		return
	}
	*/

	c.JSON(http.StatusOK, gin.H{"Success": "ok"})
}
