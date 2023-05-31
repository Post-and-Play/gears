package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

// SendMail godoc
// @Summary      Sends a mail
// @Description  With params sends a mail
// @Tags         mail
// @Accept       json
// @Produce      json
// @Param        mail  body  models.Edwiges  true  "Edwiges Model"
// @Success      200  {object}  models.Edwiges
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /mail [post]
func SendMail(c *gin.Context) {
	var edwigesRequest models.Edwiges

	if err := c.ShouldBindJSON(&edwigesRequest); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.EdwigesValidator(&edwigesRequest); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	client := infra.NewHttpClient(os.Getenv("EDWIGES_URL"))

	edwigesResponse, err := client.EdwigesPost(edwigesRequest.Mail)
	if err != nil {
		log.Default().Printf("Mail error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Mail error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": edwigesResponse})
}
