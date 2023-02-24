package controllers

import (
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no login": err.Error()})
		return
	}

	var user models.User

	infra.DB.Table("USER").Where("mail = ?", login.Mail).First(&user)

	if user.Mail != login.Mail {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no banco": "Usuário não existe!"})
		return
	}

	//Encode pass ennter && verify
	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro no login": "Credenciais invalidas!"})
		return
	}

	//Generate JWT Token
	token, err := services.NewJWTService().GenerateToken(1234)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Erro no login": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": token})
}
