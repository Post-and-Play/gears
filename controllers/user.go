package controllers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

	url := os.Getenv("FRONT_URL")

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

	if infra.DB.Where("mail = ?", user.Mail).Or("user_name = ?", user.UserName).Find(&user).RowsAffected > 0 {
		if user.ID != 0 {
			log.Default().Print("User already exists")
			c.JSON(http.StatusConflict, gin.H{"already": "E-mail ou Nickname já existente"})
			return
		}
	}

	user.Password = services.SHA256Encoder(user.Password)

	currentTime := time.Now()

	user.SecurityKey = services.SHA256Encoder(user.Mail + user.Password + currentTime.String())

	if infra.DB.Model(&user).Create(&user).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	var receiver models.Receiver
	receiver.ReceiverMail = user.Mail 

	var mailRequest models.MailRequest
	mailRequest.Subject = "PAP - Verificação do E-mail da conta"
	mailRequest.Title = "Verificar E-mail da sua conta PAP"
	mailRequest.Message =  "Olá!\n Para verificar seu e-mail da conta clique no link abaixo: "
	mailRequest.Link =  url + "/verificar-email?key=" + user.SecurityKey
	mailRequest.Footer =  "E-mail automático, não responda esse e-mail.\nEquipe Posting And Playing"
	mailRequest.ButtonText = "Verificar meu E-mail"

	services.SendMail(&receiver, &mailRequest)

	if mailRequest.OK != true {
		log.Default().Printf("Mail error: %+v", "error")
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error" : "error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// EditUser godoc
// @Summary      Edits an user
// @Description  With params edits an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  models.EditUser  true  "User Model"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /users [patch]
func EditUser(c *gin.Context) {
	id := c.Query("id")
	var edit_user models.EditUser

	if err := c.ShouldBindJSON(&edit_user); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.EditUserValidator(&edit_user); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var user models.User

	if infra.DB.Where("id = $1", id).Find(&user).RowsAffected > 0 {
		if user.ID == 0 {
			log.Default().Print("Wrong login")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}
	
	if infra.DB.Model(&user).Updates(&edit_user).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// EditPassword godoc
// @Summary      Edits EditPassword an user
// @Description  With params edits EditPassword an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  models.EditPassword  true  "User Model"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /users [put]
func EditPassword(c *gin.Context) {
	id := c.Query("id")
	var userpass models.EditPassword

	if err := c.ShouldBindJSON(&userpass); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.EditPasswordValidator(&userpass); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var user models.User

	if infra.DB.Where("id = $1", id).Find(&user).RowsAffected > 0 {
		if user.ID == 0 {
			log.Default().Print("Wrong login")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}

	userpass.Password = services.SHA256Encoder(userpass.Password)

	if infra.DB.Model(&user).Updates(&userpass).RowsAffected == 0 {
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


// SearchUsers godoc
// @Summary      Show user by filter
// @Description  Route to show games
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.User
// @Failure      404  {object}  map[string][]string
// @Router       /users/search [get]
func SearchUsers(c *gin.Context) {
	var users []models.User
	name := c.Query("name")

	if strings.Compare(name, "") != 0 { 
		//log.Default().Print("name has cotent: " + name)
		infra.DB.Where("name LIKE ? OR user_name LIKE ?", "%" + name + "%" , "%" + name + "%").Find(&users)
	} else {
		infra.DB.Find(&users)
	}

	/*if len(users) == 0 {
		log.Default().Print("No has users")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "No has users"})
		return
	} else {
		if users[0].ID == 0 {
			log.Default().Print("User not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}*/

	c.JSON(http.StatusOK, users)
}

// ListUsers godoc
// @Summary      Show users
// @Description  Route to show games
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.User
// @Failure      404  {object}  map[string][]string
// @Router       /users/list [get]
func ListUsers(c *gin.Context) {
	var users []models.User

	infra.DB.Find(&users)
		
	/*if len(users) == 0 {
		log.Default().Print("No has users")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "No has users"})
		return
	} else {
		if users[0].ID == 0 {
			log.Default().Print("User not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}*/
	

	c.JSON(http.StatusOK, users)
}
