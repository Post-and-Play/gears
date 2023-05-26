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

func ListLastReviews(c *gin.Context) {

}

func ListReviewsByUser(c *gin.Context) {

}

func DeleteReview(c *gin.Context) {

}
