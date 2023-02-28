package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

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
