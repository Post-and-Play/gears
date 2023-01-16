package controllers

import (
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
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

	infra.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	infra.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	infra.DB.First(&user, id)

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

	infra.DB.Model(&user).UpdateColumns(user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	infra.DB.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{"data": "User deleted sucessfully"})
}
