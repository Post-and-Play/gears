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

	if infra.DB.Where("followed_user_id = $1 AND following_user_id = $2", follow.FollowedUserId, follow.FollowingUserId).Find(&follow).RowsAffected > 0 {
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

	var followed models.User
	var following models.User

	if infra.DB.Where("id = $1", follow.FollowedUserId).Find(&followed).RowsAffected > 0 {
		if followed.ID == 0 {
			log.Default().Print("Wrong Followed")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Followed User not found"})
			return
		}
	}

	if infra.DB.Where("id = $1", follow.FollowingUserId).Find(&following).RowsAffected > 0 {
		if following.ID == 0 {
			log.Default().Print("Wrong Following")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Following User not found"})
			return
		}
	}
	
	if infra.DB.Model(&followed).Updates(models.User{Followed: followed.Followed + 1}).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	if infra.DB.Model(&following).Updates(models.User{Following: following.Following + 1}).RowsAffected == 0 {
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

	var followedUserId = follow.FollowedUserId
	var followingUserId = follow.FollowingUserId

	if infra.DB.Delete(&follow, id).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	var followed models.User
	var following models.User

	if infra.DB.Where("id = $1", followedUserId).Find(&followed).RowsAffected > 0 {
		if followed.ID == 0 {
			log.Default().Print("Wrong Followed")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Followed User not found"})
			return
		}
	}

	if infra.DB.Where("id = $1", followingUserId).Find(&following).RowsAffected > 0 {
		if following.ID == 0 {
			log.Default().Print("Wrong Following")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Following User not found"})
			return
		}
	}
	
	if infra.DB.Model(&followed).Updates(models.User{Followed: followed.Followed - 1}).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	if infra.DB.Model(&following).Updates(models.User{Following: following.Following - 1}).RowsAffected == 0 {
		log.Default().Print("Internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"Internal server error": "Something has occured"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OK": "Follow deleted sucessfully"})
}


// GetFollow godoc
// @Summary      Show a follow
// @Description  Route to show a follow
// @Tags         follow
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Review
// @Failure      404  {object}  map[string][]string
// @Router       /follow [get]
func GetFollow(c *gin.Context) {
	var follow models.Follow
	following_user_id := c.Query("following_user_id")
	followed_user_id := c.Query("followed_user_id")

	if infra.DB.Where("followed_user_id = $1 AND following_user_id = $2", followed_user_id, following_user_id).Find(&follow).RowsAffected > 0 {
		if follow.Id == 0 {
			log.Default().Print("Follow not found")
			c.JSON(http.StatusNotFound, gin.H{"Not found": "Follow not found"})
			return
		}
	} 

	c.JSON(http.StatusOK, follow)
}