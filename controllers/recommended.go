package controllers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"strconv"
	"github.com/Post-and-Play/gears/services"
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

// CreateRecommended godoc
// @Summary      Creates a new recommended game
// @Description  With params creates a new recommended game
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        game  body  models.Recommmended  true  "recommended Model"
// @Success      200  {object}  models.Recommmended
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /recommended [post]
func CreateRecommended(c *gin.Context) {
	var game models.Recommended

	if err := c.ShouldBindJSON(&game); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.RecommendedValidator(&game); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Where("name = $1", game.Name).Find(&game).RowsAffected > 0 {
		if game.Id != 0 {
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


// ApproveRecommended godoc
// @Summary      ApproveRecommended a new recommended game
// @Description  With params creates a new recommended game
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        game  body  models.Recommmended  true  "recommended Model"
// @Success      200  {object}  models.Recommmended
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /recommended [post]
func ApproveRecommended(c *gin.Context) {
	var game models.Recommended

	if err := c.ShouldBindJSON(&game); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.RecommendedValidator(&game); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	var recommended models.Recommended

	if infra.DB.Where("id = $1", game.Id).Find(&recommended).RowsAffected > 0 {
		if game.Id == 0 {
			log.Default().Print("Wrong login")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Recommmended not found"})
			return
		}
	}
	
	if infra.DB.Model(&recommended).Updates(&game).RowsAffected == 0 {
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
func GetRecommended(c *gin.Context) {
	var game models.Recommended
	id := c.Query("id")

	infra.DB.First(&game, id)

	if game.Id == 0 {
		log.Default().Print("Game not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// SearchGames godoc
// @Summary      Show games
// @Description  Route to show games
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Game
// @Failure      404  {object}  map[string][]string
// @Router       /games/search [get]
func SearchRecommended(c *gin.Context) {
	var games []models.Recommended

	url := os.Getenv("API_HOST")

	infra.DB.Find(&games)

	if len(games) == 0 {
		log.Default().Print("No has games")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "No has games"})
		return
	} else {
		if games[0].Id == 0 {
			log.Default().Print("Game not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Game not found"})
			return
		}
	}

	for i := 0; i < len(games); i++ {
	
		if games[i].CoverAdr != "" {
			idx := strings.Index(games[i].CoverAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("games&" + strconv.FormatUint(uint64(games[i].Id), 10) + "&cover_adr")
				games[i].CoverAdr = url + "/api/image/" + cipher
			}
		}

		if games[i].TopAdr != "" {
			idx := strings.Index(games[i].TopAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("games&" + strconv.FormatUint(uint64(games[i].Id), 10) + "&top_adr")
				games[i].TopAdr = url + "/api/image/" + cipher
			}
		}

	}

	c.JSON(http.StatusOK, games)
}

// ListGames godoc
// @Summary      Show games
// @Description  Route to show games
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Game
// @Failure      404  {object}  map[string][]string
// @Router       /games/list [get]
func ListRecommended(c *gin.Context) {
	var games []models.Recommended

	url := os.Getenv("API_HOST")

	infra.DB.Find(&games)

	if len(games) == 0 {
		log.Default().Print("No has games")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "No has games"})
		return
	} else {
		if games[0].Id == 0 {
			log.Default().Print("Game not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Game not found"})
			return
		}
	}
	
	for i := 0; i < len(games); i++ {
	
		if games[i].CoverAdr != "" {
			idx := strings.Index(games[i].CoverAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("games&" + strconv.FormatUint(uint64(games[i].Id), 10) + "&cover_adr")
				games[i].CoverAdr = url + "/api/image/" + cipher
			}
		}

		if games[i].TopAdr != "" {
			idx := strings.Index(games[i].TopAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("games&" + strconv.FormatUint(uint64(games[i].Id), 10) + "&top_adr")
				games[i].TopAdr = url + "/api/image/" + cipher
			}
		}

	}

	c.JSON(http.StatusOK, games)
}

// DeleteRecommended godoc
// @Summary      Show an user
// @Description  Route to show an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /recommended [delete]
func DeleteRecommended(c *gin.Context) {
	var recommended models.Recommended
	id := c.Query("id")

	infra.DB.First(&recommended, id)
	if recommended.Id == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
	}

	if infra.DB.Delete(&recommended, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Recommended deleted sucessfully"})
}
