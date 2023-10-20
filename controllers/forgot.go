package controllers

import (
	"log"
	"net/http"
	"os"
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
	url := os.Getenv("FRONT_URL")

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

	var receiver models.Receiver
	receiver.ReceiverMail = email 

	var mailRequest models.MailRequest
	mailRequest.Subject = "PAP - Redefinição de senha"
	mailRequest.Title = "Redefinir a senha da sua conta PAP"
	mailRequest.Message =  "Olá Para redefinir sua senha clique no link abaixo: "
	mailRequest.Link =  url + "/redefinir-senha?key=" + user.SecurityKey
	mailRequest.Footer =  "E-mail automático, não responda esse e-mail.\nEquipe Posting And Playing"
	mailRequest.ButtonText = "Redefinir minha senha"

	services.SendMail(&receiver, &mailRequest)

	if mailRequest.OK != true {
		log.Default().Printf("Mail error: %+v", "error")
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error" : "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": mailRequest.OK})
}

// GetRecoverUser godoc
// @Summary      Get recover account
// @Description  Route to get key for recover
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string][]string
// @Router       /recover [get]
func GetRecoverUser(c *gin.Context) {
	var user models.User
	key := c.Query("key")

	if infra.DB.Model(&user).Select("id").Where("security_key = ?", key).Find(&user).RowsAffected == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	if user.ID == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID})
}

// RecoverPasswordUser godoc
// @Summary      Recover password
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.RecoverPasswordUser  true  "RecoverPasswordUser Model"
// @Success      200  {object}  models.Edwiges
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /recover [post]
func RecoverPasswordUser(c *gin.Context) {
	id := c.Query("id")

	var recover models.RecoverPasswordUser
	var user models.User

	if err := c.ShouldBindJSON(&recover); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.RecoverPasswordUserValidator(&recover); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	log.Default().Print("id: " + id + "\n" + "security_key: " + recover.SecurityKey)

	if infra.DB.Where("id = ? AND security_key = ?", id, recover.SecurityKey).Find(&user).RowsAffected > 0 {
		if user.ID == 0 {
			log.Default().Print("User not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}

	recover.Password = services.SHA256Encoder(recover.Password)
	recover.SecurityKey = ""

	if infra.DB.Model(&user).Updates(&recover).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "ok"})
}



// ForgotAdmin godoc
// @Summary      ForgotAdmin password
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.Forgot  true  "Forgot Model"
// @Success      200  {object}  models.Forgot
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /admins/forgot [get]
func ForgotAdmin(c *gin.Context) {
	var admin models.Admin
	email := c.Query("mail")
	url := os.Getenv("FRONT_URL")

	infra.DB.Where("mail = ?", email).First(&admin)

	if admin.ID == 0 {
		log.Default().Print("admin not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "admin not found"})
		return
	}

	currentTime := time.Now()

	admin.SecurityKey = services.SHA256Encoder(admin.Mail + admin.Password + currentTime.String())

	if infra.DB.Model(&admin).Updates(&admin).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	var receiver models.Receiver
	receiver.ReceiverMail = email 

	var mailRequest models.MailRequest
	mailRequest.Subject = "PAP - Redefinição de senha"
	mailRequest.Title = "Redefinir a senha da sua conta PAP"
	mailRequest.Message =  "Olá Para redefinir sua senha clique no link abaixo: "
	mailRequest.Link =  url + "/admin/redefinir-senha?key=" + admin.SecurityKey
	mailRequest.Footer =  "E-mail automático, não responda esse e-mail.\nEquipe Posting And Playing"
	mailRequest.ButtonText = "Redefinir minha senha"

	services.SendMail(&receiver, &mailRequest)

	if mailRequest.OK != true {
		log.Default().Printf("Mail error: %+v", "error")
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error" : "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": mailRequest.OK})
}

// GetRecoverAdmin godoc
// @Summary      Get recover account
// @Description  Route to get key for recover
// @Tags         admins
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Admin
// @Failure      404  {object}  map[string][]string
// @Router       /admins/recover [get]
func GetRecoverAdmin(c *gin.Context) {
	var admin models.Admin
	key := c.Query("key")

	if infra.DB.Model(&admin).Select("id").Where("security_key = ?", key).Find(&admin).RowsAffected == 0 {
		log.Default().Print("admin not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "admin not found"})
		return
	}

	if admin.ID == 0 {
		log.Default().Print("admin not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": admin.ID})
}

// RecoverPasswordUser godoc
// @Summary      Recover password
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.RecoverPasswordAdmin  true  "RecoverPasswordAdmin Model"
// @Success      200  {object}  models.Edwiges
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /admins/recover [post]
func RecoverPasswordAdmin(c *gin.Context) {
	id := c.Query("id")

	var recover models.RecoverPasswordAdmin
	var admin models.Admin

	if err := c.ShouldBindJSON(&recover); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.RecoverPasswordAdminValidator(&recover); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	log.Default().Print("id: " + id + "\n" + "security_key: " + recover.SecurityKey)

	if infra.DB.Where("id = ? AND security_key = ?", id, recover.SecurityKey).Find(&admin).RowsAffected > 0 {
		if admin.ID == 0 {
			log.Default().Print("admin not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "admin not found"})
			return
		}
	}

	recover.Password = services.SHA256Encoder(recover.Password)
	recover.SecurityKey = ""

	if infra.DB.Model(&admin).Updates(&recover).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "ok"})
}
