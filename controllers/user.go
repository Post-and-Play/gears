package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("\nUSER 1: %+v", user)

	if err := models.UserValidator(&user); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	fmt.Printf("\nUSER 2: %+v", user)

	if infra.DB.First(&user, user.Mail).RowsAffected > 0 {
		if user.ID != "" {
			log.Default().Print("User already exists")
			c.JSON(http.StatusConflict, user)
			return
		}
	}

	fmt.Printf("\nUSER 3: %+v", user)

	user.Password = services.SHA256Encoder(user.Password)

	if infra.DB.Create(&user).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		fmt.Println("ERRO AQUIIII")
		fmt.Println(user)
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	fmt.Printf("\nUSER 4: %+v", user)

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")

	infra.DB.First(&user, id)

	if user.ID == "" {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Binding error": err.Error()})
		return
	}

	if err := models.UserValidator(&user); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var databaseUser models.User

	infra.DB.First(&databaseUser, user.Mail)
	if user.ID == "" {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	if infra.DB.Model(&user).UpdateColumns(&user).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")

	infra.DB.First(&user, id)
	if user.ID == "" {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
	}

	if infra.DB.Delete(&user, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "User deleted sucessfully"})
}
