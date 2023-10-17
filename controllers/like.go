package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

// LikeReview godoc
// @Summary      Likes a review
// @Description  With params likes a review
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        like  body  models.Like  true  "Like Model"
// @Success      200  {object}  models.Like
// @Failure      400  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /like [post]
func LikeReview(c *gin.Context) {
	var like models.Like

	if err := c.ShouldBindJSON(&like); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.LikeValidator(&like); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Model(&like).Create(&like).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	var review models.Review

	if infra.DB.Model(&review).Where("id = ?", like.ReviewId).Update("likes", review.Likes + 1).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, like)
}

// GetLikesByUser godoc
// @Summary      Show likes by user
// @Description  Route to show likes by user
// @Tags         likes
// @Accept       json
// @Produce      json
// @Success      200  {object}  int
// @Failure      404  {object}  map[string][]string
// @Router       /likes/user [get]
func GetLikesByUser(c *gin.Context) {
	var likes []models.Like

	id := c.Query("id")

	infra.DB.Find(&likes).Where("user_id = $1", id)

	if len(likes) == 0 {
		log.Default().Print("No has likes")
		c.JSON(http.StatusOK, likes)
		return
	}

	c.JSON(http.StatusOK, likes)
}

// GetLikesByReview godoc
// @Summary      Show likes by review
// @Description  Route to show likes by review
// @Tags         likes
// @Accept       json
// @Produce      json
// @Success      200  {object}  int
// @Failure      404  {object}  map[string][]string
// @Router       /likes/review [get]
func GetLikesByReview(c *gin.Context) {
	var likes int64

	id := c.Query("id")

	infra.DB.Count(&likes).Where("review_id = $1", id)

	c.JSON(http.StatusOK, likes)
}

// UnlikeReview godoc
// @Summary      Show a like
// @Description  Route to delete a like
// @Tags         likes
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /like [delete]
func UnlikeReview(c *gin.Context) {
	var like models.Like
	user_id := c.Query("user_id")
	review_id := c.Query("review_id")

	//infra.DB.Find(&like).Where("user_id = $1 AND review_id = $1", user_id, review_id)

	//if like.Id == 0 {
	//	log.Default().Print("like not found")
	//	c.JSON(http.StatusNotFound, gin.H{"Not found": "Like not found"})
	//}

	if infra.DB.Where("user_id = $1 AND review_id = $2", user_id, review_id).Delete(&like).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Like deleted sucessfully"})
}
