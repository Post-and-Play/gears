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
		"/healthz",
		http.MethodGet,
		controllers.Health,
	},
	{
		"/readiness",
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
		http.MethodPatch,
		controllers.EditUser,
	},
	{
		"/users",
		http.MethodPut,
		controllers.EditPassword,
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
		"/games",
		http.MethodGet,
		controllers.GetGame,
	},
	{
		"/games/search",
		http.MethodGet,
		controllers.SearchGames,
	},
	{
		"/games/ranking",
		http.MethodGet,
		controllers.GetRanking,
	},
	{
		"/games",
		http.MethodPost,
		controllers.CreateGame,
	},
}

var review = []Route{
	{
		"/review",
		http.MethodPost,
		controllers.CreateReview,
	},
	{
		"/review",
		http.MethodDelete,
		controllers.DeleteReview,
	},
	{
		"/review",
		http.MethodGet,
		controllers.GetReview,
	},
	{
		"/reviews",
		http.MethodGet,
		controllers.ListLastReviews,
	},
	{
		"/reviews/user",
		http.MethodGet,
		controllers.ListReviewsByUser,
	},
	{
		"/reviews/game",
		http.MethodGet,
		controllers.ListReviewsByGame,
	},
}

var like = []Route{
	{
		"/like",
		http.MethodPost,
		controllers.LikeReview,
	},
	{
		"/like",
		http.MethodDelete,
		controllers.UnlikeReview,
	},
	{
		"/likes/user",
		http.MethodGet,
		controllers.GetLikesByUser,
	},
	{
		"/likes/review",
		http.MethodGet,
		controllers.GetLikesByReview,
	},
}

var follow = []Route{
	{
		"/follow",
		http.MethodPost,
		controllers.Follow,
	},
	{
		"/unfollow",
		http.MethodDelete,
		controllers.Unfollow,
	},
}

var mail = []Route{
	{
		"/mail",
		http.MethodPost,
		controllers.SendMail,
	},
}
