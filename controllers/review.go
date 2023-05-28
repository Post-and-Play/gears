package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
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

	if review.ID == 0 {
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
	var reviews []models.Review

	infra.DB.Find(reviews).Limit(30)

	if reviews[0].ID == 0 {
		log.Default().Print("Reviews not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Reviews not found"})
		return
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

	infra.DB.Find(reviews).Where("user_id = $1", id).Limit(30)

	if reviews[0].ID == 0 {
		log.Default().Print("Reviews not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Reviews not found"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// DeleteReview godoc
// @Summary      Show an review
// @Description  Route to show an review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /reviews [delete]
func DeleteReview(c *gin.Context) {
	var review models.Review
	id := c.Query("id")

	infra.DB.First(&review, id)
	if review.ID == 0 {
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
