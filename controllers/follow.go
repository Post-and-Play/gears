package controllers

import (
	"log"
	"net/http"

	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

// Follow godoc
// @Summary      Follow a user
// @Description  With params follows a user
// @Tags         follow
// @Accept       json
// @Produce      json
// @Param        follow  body  models.Follow  true  "Follow Model"
// @Success      200  {object}  models.Follow
// @Failure      400  {object}  map[string][]string
// @Failure      409  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /follow [post]
func Follow(c *gin.Context) {
	var follow models.Follow

	if err := c.ShouldBindJSON(&follow); err != nil {
		log.Default().Printf("Binding error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.FollowValidator(&follow); err != nil {
		log.Default().Printf("Validation error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Validation error": err.Error()})
		return
	}

	if infra.DB.Where("followed_id = $1 AND following_id = $2", follow.FollowedUserId, follow.FollowingUserId).Find(&follow).RowsAffected > 0 {
		if follow.Id != 0 {
			log.Default().Print("Follow already exists")
			c.JSON(http.StatusConflict, follow)
			return
		}
	}

	if infra.DB.Model(&follow).Create(&follow).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, follow)
}

// Unfollow godoc
// @Summary      Unfollow a user
// @Description  Route to unfollow a user
// @Tags         follow
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  map[string][]string
// @Failure      404  {object}  map[string][]string
// @Failure      500  {object}  map[string][]string
// @Router       /follow [delete]
func Unfollow(c *gin.Context) {
	var follow models.Follow
	id := c.Query("id")

	infra.DB.First(&follow, id)
	if follow.Id == 0 {
		log.Default().Print("Follow not found")
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Follow not found"})
	}

	if infra.DB.Delete(&follow, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Follow deleted sucessfully"})
}
