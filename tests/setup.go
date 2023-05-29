package tests

import (
	"github.com/Post-and-Play/gears/infra"
	"github.com/Post-and-Play/gears/models"
	"github.com/gin-gonic/gin"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func UserMock() {
	user := models.User{ID: 1, Name: "Jonas", UserName: "john_gamer", Password: "churros", Mail: "jonas_brothers", BirthDate: "01/02/2002"}
	infra.DB.Create(user)
}

func DeleteUserMock() {
	var user models.User
	infra.DB.Delete(&user, ID)
}

func GameMock() {
	game := models.Game{Id: 1, Name: "r6", Genders: "sasuke-naruto", CoverAdr: "xpto", Rating: 3.2}
	infra.DB.Create(game)
}

func DeleteGameMock() {
	var game models.Game
	infra.DB.Delete(&game, ID)
}
