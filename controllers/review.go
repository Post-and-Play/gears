package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/Post-and-Play/gears/services"
	"github.com/gin-gonic/gin"
)

// CreateReview godoc
// @Summary      Creates a new review
// @Description  With params creates a new review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        review  body  models.Review  true  "Review Model"
// @Success      200  {object}  models.Review
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /review [post]
func CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ReviewValidator(&review); err != nil {
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

// GetReview godoc
// @Summary      Show a review
// @Description  Route to show a review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Review
// @Failure      404  {object}  map[string][]string
// @Router       /review [get]
func GetReview(c *gin.Context) {
	var review models.Review
	id := c.Query("id")

	infra.DB.First(&review, id)

	if review.Id == 0 {
		log.Default().Print("Review not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Review not found"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// ListLastReviews godoc
// @Summary      Show last reviews
// @Description  Route to show last reviews
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Review
// @Failure      404  {object}  map[string][]string
// @Router       /reviews [get]
func ListLastReviews(c *gin.Context) {
	var reviews []models.ReviewUser
	var review  []models.Review

	url := os.Getenv("API_HOST")

	infra.DB.Model(&review).Select("reviews.id, reviews.user_id, reviews.game_id, reviews.grade, reviews.image_adr, reviews.opinion, reviews.likes, users.name, users.photo_adr, games.name AS game_name, games.top_adr").Joins("LEFT JOIN users ON users.id = reviews.user_id ").Joins("LEFT JOIN games ON games.id = reviews.game_id ").Scan(&reviews).Limit(30)

	for i := 0; i < len(reviews); i++ {
		//fmt.Println(i)
		if reviews[i].ImageAdr != "" {
			cipher := Encrypt("m=reviews&uid=" + strconv.FormatUint(uint64(reviews[i].Id), 10) + "&att=image_adr")
			reviews[i].ImageAdr = url + "/api/image?" + cipher
		}

		if reviews[i].PhotoAdr != "" {
			cipher := Encrypt("m=users&uid=" + strconv.FormatUint(uint64(reviews[i].UserId), 10) + "&att=photo_adr")
			reviews[i].PhotoAdr = url + "/api/image?" + cipher
		}

		if reviews[i].TopAdr != "" {
			cipher := Encrypt("m=games&uid=" + strconv.FormatUint(uint64(reviews[i].GameId), 10) + "&att=top_adr")
			reviews[i].TopAdr = url + "/api/image?" + cipher
		}

	}

	c.JSON(http.StatusOK, reviews)
}

// ListReviewsByUser godoc
// @Summary      Show last reviews by user
// @Description  Route to show last reviews by user
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Review
// @Failure      404  {object}  map[string][]string
// @Router       /reviews/user [get]
func ListReviewsByUser(c *gin.Context) {
	var reviews []models.Review

	id := c.Query("id")

	infra.DB.Find(&reviews).Where("user_id = $1", id).Limit(30)

    //if reviews[0].Id == 0 {
	//	log.Default().Print("Reviews not found")
	//	c.JSON(http.StatusNotFound, gin.H{"Not found": "Reviews not found"})
	//	return
	//}


	c.JSON(http.StatusOK, reviews)
}

// DeleteReview godoc
// @Summary      Delete a review
// @Description  Route to delete a review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /review [delete]
func DeleteReview(c *gin.Context) {
	var review models.Review
	id := c.Query("id")

	infra.DB.First(&review, id)
	if review.Id == 0 {
		log.Default().Print("Review not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Review not found"})
	}

	if infra.DB.Delete(&review, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Review deleted sucessfully"})
}


// ListReviewsByGame godoc
// @Summary      Show last reviews by user
// @Description  Route to show last reviews by user
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Review
// @Failure      404  {object}  map[string][]string
// @Router       /reviews/game [get]
func ListReviewsByGame(c *gin.Context) {
	var reviews []models.ReviewUser
	var review  []models.Review

	id := c.Query("id")

	infra.DB.Model(&review).Select("reviews.id, reviews.user_id, reviews.game_id, reviews.grade, reviews.image_adr, reviews.opinion, reviews.likes, users.name, users.photo_adr, games.name AS game_name, games.top_adr").Joins("LEFT JOIN users ON users.id = reviews.user_id").Joins("LEFT JOIN games ON games.id = reviews.game_id ").Where("game_id = $1", id).Scan(&reviews).Limit(30)

    //if reviews[0].Id == 0 {
	//	log.Default().Print("Reviews not found")
	//	c.JSON(http.StatusNotFound, gin.H{"Not found": "Reviews not found"})
	//	return
	//}

	c.JSON(http.StatusOK, reviews)
}