package controllers

import (
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.UserValidator(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	infra.DB.First(&user).Where("mail = $1", user.Mail)
	if user.ID != "" {
		c.JSON(http.StatusOK, user)
	}

	user.Password = services.SHA256Encoder(user.Password)

	infra.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")

	infra.DB.First(&user, id)

	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.UserValidator(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var databaseUser models.User

	infra.DB.First(&databaseUser).Where("mail = $1", databaseUser.Mail)
	if user.ID != "" {
		c.JSON(http.StatusOK, user)
	}

	infra.DB.Model(&user).UpdateColumns(user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")

	infra.DB.First(&user, id)
	if id == "" {
		c.JSON(http.StatusConflict, gin.H{"Conflict": "User not exist"})
	}

	infra.DB.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{"OK": "User deleted sucessfully"})
}
