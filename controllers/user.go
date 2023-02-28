package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// CreateUser godoc
// @Summary      Creates a new user
// @Description  With params creates a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Model"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UserValidator(&user); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Where("mail = $1", user.Mail).Find(&user).RowsAffected > 0 {
		if user.ID != 0 {
			log.Default().Print("User already exists")
			c.JSON(http.StatusConflict, user)
			return
		}
	}

	user.Password = services.SHA256Encoder(user.Password)

	if infra.DB.Model(&user).Create(&user).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser godoc
// @Summary      Show an user
// @Description  Route to show an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string][]string
// @Router       /users [get]
func GetUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")

	infra.DB.First(&user, id)

	if user.ID == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Show an user
// @Description  Route to show an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /users [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")

	infra.DB.First(&user, id)
	if user.ID == 0 {
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
