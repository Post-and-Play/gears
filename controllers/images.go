package controllers

import (
	"log"
	"net/http"
    "strings"
    "encoding/base64"
	"github.com/Post-and-Play/gears/infra"
    "github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

// GetImage godoc
// @Summary      Show a image in database
// @Description  Route to show a game
// @Tags         images
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Game
// @Failure      404  {object}  map[string][]string
// @Router       /image [get]
func GetImage(c *gin.Context) {

	var img64 string
    uid := c.Param("id")
    text := services.Decrypt(uid)

    log.Println("Hash %x", uid)
    log.Println("Text %x", text)

    itens := strings.Split(text, "&")

    if len(itens) < 3 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Not content"})
        return
    }

	m := itens[0]
	id := itens[1]
	att := itens[2]

	infra.DB.Table(m).Select(att).Where("id = ?", id).Find(&img64)

    if img64 == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not content"})
        return
    }

	idx := strings.Index(img64, ";base64,")
    if idx < 0 {
        panic("InvalidImage")
    }

    ImageType := img64[11:idx]
    log.Println(ImageType)
    unbased, err := base64.StdEncoding.DecodeString(img64[idx+8:])
    if err != nil {
        panic("Cannot decode b64")
    }

    switch ImageType {
    case "png":
     
        c.Data(http.StatusOK, "image/png", unbased)

    case "jpeg":

        c.Data(http.StatusOK, "image/jpeg", unbased)

    case "gif":
     
        c.Data(http.StatusOK, "image/gif", unbased)

    case "x-icon":
     
        c.Data(http.StatusOK, "image/x-icon", unbased)

    case "webp":
     
        c.Data(http.StatusOK, "image/webp", unbased)

    }

	//bin := base64toJpg(img, m + id + att)

	//c.JSON(http.StatusOK, image)

}
