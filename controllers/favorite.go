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
)

// FavoriteGame godoc
// @Summary      Favorite a game
// @Description  With params Favorite a game
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Param        like  body  models.Favorite  true  "Favorite Model"
// @Success      200  {object}  models.Favorite
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /favorite [post]
func FavoriteGame(c *gin.Context) {
	var favorite models.Favorite

	if err := c.ShouldBindJSON(&favorite); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.FavoriteValidator(&favorite); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Model(&favorite).Create(&favorite).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, favorite)
}

// GetFavoritesByUser godoc
// @Summary      Show Favorites by user
// @Description  Route to show favorite by user
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Success      200  {object}  int
// @Failure      404  {object}  map[string][]string
// @Router       /favorite/user [get]
func GetFavoritesByUser(c *gin.Context) {
	var favorites []models.FavoriteGame
	var favorite models.Favorite

	url := os.Getenv("API_HOST")

	id := c.Query("id")

	infra.DB.Model(&favorite).Select("favorites.id, favorites.user_id, favorites.game_id, games.name, games.genders, games.description, games.cover_adr, games.top_adr").Joins("LEFT JOIN games ON games.id = favorites.game_id").Where("user_id = $1", id).Scan(&favorites)

	if len(favorites) == 0 {
		log.Default().Print("No has favorites")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "No has favorites"})
		return
	} else {
		if favorites[0].Id == 0 {
			log.Default().Print("Favorites not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Favorites not found"})
			return
		}
	}

	for i := 0; i < len(favorites); i++ {
	
		if favorites[i].CoverAdr != "" {
			idx := strings.Index(favorites[i].CoverAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("games&" + strconv.FormatUint(uint64(favorites[i].GameId), 10) + "&cover_adr")
				favorites[i].CoverAdr = url + "/api/image/" + cipher
			}
		}

		if favorites[i].TopAdr != "" {
			idx := strings.Index(favorites[i].TopAdr, ";base64,")
			if idx >= 0 {
				cipher := services.Encrypt("games&" + strconv.FormatUint(uint64(favorites[i].GameId), 10) + "&top_adr")
				favorites[i].TopAdr = url + "/api/image/" + cipher
			}
		}

	}

	c.JSON(http.StatusOK, favorites)
}

// GetFavoritesByGame godoc
// @Summary      Show favorites by review
// @Description  Route to show favorites by review
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Success      200  {object}  int
// @Failure      404  {object}  map[string][]string
// @Router       /favorite/game [get]
func GetFavoritesByGame(c *gin.Context) {
	var favorite []models.Favorite
	var count int64 = 0
	id := c.Query("id")

	infra.DB.Model(&favorite).Where("game_id = $1", id).Count(&count)

	if count == 0 {
		log.Default().Print("No has favorites")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "No has favorites"})
		return
	} else {
		if favorite[0].Id == 0 {
			log.Default().Print("Favorites not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Favorites not found"})
			return
		}
	}

	c.JSON(http.StatusOK, count)
}

// UnFavoriteGame godoc
// @Summary      Remove favorite 
// @Description  Route to delete a favorite
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /favorite [delete]
func UnFavoriteGame(c *gin.Context) {
	var favorite models.Favorite
	id := c.Query("id")

	infra.DB.First(&favorite, id)
	if favorite.Id == 0 {
		log.Default().Print("favorite not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "favorite not found"})
	}

	if infra.DB.Delete(&favorite, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Favorite deleted sucessfully"})
}
