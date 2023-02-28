package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// GetGame godoc
// @Summary      Show a game
// @Description  Route to show a game
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string][]string
// @Router       /games [get]
func GetGame(c *gin.Context) {
	var game models.Game
	id := c.Query("id")

	infra.DB.First(&game, id)

	if game.ID == 0 {
		log.Default().Print("Game not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}
