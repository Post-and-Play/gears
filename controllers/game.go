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

// CreateGame godoc
// @Summary      Creates a new game
// @Description  With params creates a new game
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        game  body  models.Game  true  "Game Model"
// @Success      200  {object}  models.Game
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Router       /games [post]
func CreateGame(c *gin.Context) {
	var game models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.GameValidator(&game); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Where("name = $1", game.Name).Find(&game).RowsAffected > 0 {
		if game.ID != 0 {
			log.Default().Print("Game already exists")
			c.JSON(http.StatusConflict, game)
			return
		}
	}

	if infra.DB.Model(&game).Create(&game).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// GetGame godoc
// @Summary      Show a game
// @Description  Route to show a game
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Game
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
