package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      login
// @Description  With params login
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        login  body  models.Login  true  "Login Model"
// @Success      200  {object}  models.Login
// @Failure      400  {object}  map[string][]string
// @Failure      403  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Router       /login [post]
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

	if infra.DB.Where("mail = $1", login.Login).Find(&user).RowsAffected > 0 {
		if user.ID == 0 {
			if infra.DB.Where("user_name = $1", login.Login).Find(&user).RowsAffected > 0 {
				if user.ID == 0 {
					log.Default().Print("Wrong login")
					c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
					return
				}
			}
		}
	}

	//Encode pass ennter && verify
	if user.Password != services.SHA256Encoder(login.Password) {
		log.Default().Print("Wrong password")
		c.JSON(http.StatusForbidden, gin.H{"Forbidden": "Invalid credentials"})
		return
	}

	//Generate JWT Token
	token, err := services.NewJWTService().GenerateToken(int64(user.ID))
	if err != nil {
		log.Default().Printf("Generate token error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
}
