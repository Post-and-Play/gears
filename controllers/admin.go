package controllers

import (
	"log"
	"net/http"
	"strings"
	"os"
	"strconv"
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// CreateUser godoc
// @Summary      Creates a new admin user
// @Description  With params creates a new admin user
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        user  body  models.Admin  true  "Admin Model"
// @Success      200  {object}  models.Admin
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /admins [post]

func CreateAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.AdminValidator(&admin); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Where("mail = $1", admin.Mail).Find(&admin).RowsAffected > 0 {
		if admin.ID != 0 {
			log.Default().Print("User already exists")
			c.JSON(http.StatusConflict, admin)
			return
		}
	}

	admin.Password = services.SHA256Encoder(admin.Password)

	if infra.DB.Model(&admin).Create(&admin).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// EditAdmin godoc
// @Summary      Edits an admin
// @Description  With params edits an admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        user  body  models.EditAdmin  true  "Admin Model"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /admins [patch]

func EditAdmin(c *gin.Context) {
	id := c.Query("id")
	var edit_admin models.EditAdmin

	if err := c.ShouldBindJSON(&edit_admin); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.EditAdminValidator(&edit_admin); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var admin models.Admin

	if infra.DB.Where("id = $1", id).Find(&admin).RowsAffected > 0 {
		if admin.ID == 0 {
			log.Default().Print("Wrong login")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}
	
	if infra.DB.Model(&admin).Updates(&edit_admin).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// EditPassword godoc
// @Summary      Edits EditPassword an admin
// @Description  With params edits EditPassword an admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        user  body  models.EditAdminPassword  true  "User Model"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /admins [put]

func EditAdminPassword(c *gin.Context) {
	id := c.Query("id")
	var adminpass models.EditAdminPassword

	if err := c.ShouldBindJSON(&adminpass); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.EditAdminPasswordValidator(&adminpass); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var admin models.Admin

	if infra.DB.Where("id = ?", id).Find(&admin).RowsAffected > 0 {
		if admin.ID == 0 {
			log.Default().Print("Wrong login")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
			return
		}
	}

	adminpass.Password = services.SHA256Encoder(adminpass.Password)

	if infra.DB.Model(&admin).Updates(&adminpass).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, admin)
}


// GetAdmin godoc
// @Summary      Show an admin
// @Description  Route to show an admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Admin
// @Failure      404  {object}  map[string][]string
// @Router       /admins [get]

func GetAdmin(c *gin.Context) {
	var admin models.Admin
	id := c.Query("id")

	infra.DB.First(&admin, id)

	if admin.ID == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}


// DeleteAdmin godoc
// @Summary      Show an admin
// @Description  Route to show an admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /admins [delete]
func DeleteAdmin(c *gin.Context) {
	var user models.Admin
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
// @Summary      Show admins by filter
// @Description  Route to show admins
// @Tags         admins
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Admin
// @Failure      404  {object}  map[string][]string
// @Router       /admins/search [get]
func SearchAdmins(c *gin.Context) {
	var admins []models.Admin

	url := os.Getenv("API_HOST")

	name := c.Query("name")

	if strings.Compare(name, "") != 0 { 
		//log.Default().Print("name has cotent: " + name)
		infra.DB.Where("name LIKE ?", "%" + name + "%").Find(&admins)
	} else {
		infra.DB.Find(&admins)
	}

	for i := 0; i < len(admins); i++ {
	
		if admins[i].PhotoAdr != "" {
			idx := strings.Index(admins[i].PhotoAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("admins&" + strconv.FormatUint(uint64(admins[i].ID), 10) + "&photo_adr")
				admins[i].PhotoAdr = url + "/api/image/" + cipher
			}
		}

	}

	c.JSON(http.StatusOK, admins)
}

// ListAdmins godoc
// @Summary      Show admins
// @Description  Route to show admins
// @Tags         admins
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Admin
// @Failure      404  {object}  map[string][]string
// @Router       /admins/list [get]
func ListAdmins(c *gin.Context) {
	var admins []models.Admin

	url := os.Getenv("API_HOST")

	infra.DB.Find(&admins)

	for i := 0; i < len(admins); i++ {
	
		if admins[i].PhotoAdr != "" {
			idx := strings.Index(admins[i].PhotoAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("admins&" + strconv.FormatUint(uint64(admins[i].ID), 10) + "&photo_adr")
				admins[i].PhotoAdr = url + "/api/image/" + cipher
			}
		}

	}

	c.JSON(http.StatusOK, admins)
}