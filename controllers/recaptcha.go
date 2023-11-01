package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

// VerifyRecaptcha godoc
// @Summary      Verify recaptcha
// @Description  With params verify
// @Tags         recaptcha
// @Accept       json
// @Produce      json
// @Param        review  body  models.Recaptcha  true  "Recaptcha Model"
// @Success      200  {object}  models.Recaptcha
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /recaptcha [post]
func VerifyRecaptcha(c *gin.Context) {
	var review models.Recaptcha

	if err := c.ShouldBindJSON(&review); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.RecaptchaValidator(&review); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Model(&review).Create(&review).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, review)
}