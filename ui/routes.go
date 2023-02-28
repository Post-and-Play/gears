package ui

import (
	"net/http"

	"github.com/Post-and-Play/gears/controllers"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

var healthCheck = []Route{
	{
		"healthz",
		http.MethodGet,
		controllers.Health,
	},
	{
		"readiness",
		http.MethodGet,
		controllers.Readiness,
	},
}

var swagg = []Route{
	{
		"/swagger/*any",
		http.MethodGet,
		ginSwagger.WrapHandler(swaggerfiles.Handler),
	},
}

var cad = []Route{
	{
		"/users",
		http.MethodGet,
		controllers.GetUser,
	},
	{
		"/users",
		http.MethodPost,
		controllers.CreateUser,
	},
	{
		"/users",
		http.MethodDelete,
		controllers.DeleteUser,
	},
}

var login = []Route{
	{
		"/login",
		http.MethodPost,
		controllers.Login,
	},
}

var game = []Route{
	{
		"/game",
		http.MethodGet,
		controllers.GetGame,
	},
}
