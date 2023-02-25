package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		log.Default().Printf("Biding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Binding error": err.Error()})
		return
	}

	if err := models.LoginValidator(&login); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var user models.User

	infra.DB.Table("USER").First(&user, login.Mail)

	if user.ID == "" {
		log.Default().Print("Wrong e-mail")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	//Encode pass ennter && verify
	if user.Password != services.SHA256Encoder(login.Password) {
		log.Default().Print("Wrong password")
		c.JSON(http.StatusForbidden, gin.H{"Forbidden": "Invalid credentials"})
		return
	}

	numId, err := strconv.Atoi(user.ID)
	if err != nil {
		log.Default().Printf("Strconv error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Not found": err.Error()})
		return
	}

	//Generate JWT Token
	token, err := services.NewJWTService().GenerateToken(int64(numId))
	if err != nil {
		log.Default().Printf("Generate token error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
}
