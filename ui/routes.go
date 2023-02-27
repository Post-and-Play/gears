package ui

import (
	"net/http"

	"github.com/Post-and-Play/gears/controllers"
	"github.com/gin-gonic/gin"
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
