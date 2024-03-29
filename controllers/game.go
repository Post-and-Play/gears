package controllers

import (
	"log"
	"net/http"
	"strings"
	"os"
	"strconv"
	"github.com/Post-and-Play/gears/services"
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
// @Failure      500  {object}  map[string][]string
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


// UpdateGame godoc
// @Summary      Update a game
// @Description  With params updates a game
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        game  body  models.Game  true  "Game Model"
// @Success      200  {object}  models.Game
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /games [put]
func UpdateGame(c *gin.Context) {
	var game models.Game
	var edit_game models.EditGame

	id := c.Query("id")

	infra.DB.First(&game, id)

	if game.Id == 0 {
		log.Default().Print("Game not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Game not found"})
		return
	}

	if err := c.ShouldBindJSON(&edit_game); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.EditGameValidator(&edit_game); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Model(&game).Updates(&edit_game).RowsAffected == 0 {
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
func SearchGames(c *gin.Context) {
	var games []models.Game
	name := c.Query("name")

	url := os.Getenv("API_HOST")


	if strings.Compare(name, "") != 0 { 
		//log.Default().Print("name has cotent: " + name)
		infra.DB.Where("name LIKE ?", "%" + name + "%").Find(&games)
	} else {
		infra.DB.Find(&games)
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
func ListGames(c *gin.Context) {
	var games []models.Game

	url := os.Getenv("API_HOST")

	infra.DB.Find(&games)

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

// GetRanking godoc
// @Summary Show a ranking
// @Description Route to show a ranking
// @Tags games
// @Accept       json
// @Produce      json
// @Success      200  {object}  int
// @Failure      404  {object}  map[string][]string
// @Router       /games/ranking [get]
func GetRanking(c *gin.Context) {
	var games []models.Game
	
	url := os.Getenv("API_HOST")

	infra.DB.Model(&games).Order("reviews desc").Scan(&games).Limit(5)

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


// GetSimilar godoc
// @Summary Show a similar games
// @Description Route to show a ranking
// @Tags games
// @Accept       json
// @Produce      json
// @Success      200  {object}  int
// @Failure      404  {object}  map[string][]string
// @Router       /games/similar [get]
func GetSimilar(c *gin.Context) {
	var games []models.Game
	gender := c.Query("gender")

	url := os.Getenv("API_HOST")

	infra.DB.Model(&games).Where("genders LIKE ?", "%" + gender + "%").Find(&games).Limit(5)

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
func DeleteGame(c *gin.Context) {
	var game models.Game
	id := c.Query("id")

	infra.DB.First(&game, id)
	if game.Id == 0 {
		log.Default().Print("User not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
	}

	if infra.DB.Delete(&game, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Game deleted sucessfully"})
}