package views

import (
	"github.com/Post-and-Play/gears/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.GET("/users?id", controllers.GetUser)

	r.POST("/users", controllers.CreateUser)

	r.PATCH("/users", controllers.EditUser)

	r.DELETE("/users?id", controllers.DeleteUser)

	r.POST("/login", controllers.Login)

	r.Run()
}
