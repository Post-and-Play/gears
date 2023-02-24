package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		log.Panicf("Biding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Binding error": err.Error()})
		return
	}

	var user models.User

	infra.DB.Table("USER").Where("mail = ?", login.Mail).First(&user)

	if user.ID == "" {
		log.Panic("Wrong e-mail")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Usuário não existe!"})
		return
	}

	//Encode pass ennter && verify
	if user.Password != services.SHA256Encoder(login.Password) {
		log.Panic("Wrong password")
		c.JSON(http.StatusBadRequest, gin.H{"Erro no login": "Credenciais invalidas!"})
		return
	}

	//Generate JWT Token
	token, err := services.NewJWTService().GenerateToken(1234)
	if err != nil {
		log.Panicf("Generate token error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Erro no login": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
}
