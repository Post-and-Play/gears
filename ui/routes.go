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
	{
		"/users/search",
		http.MethodGet,
		controllers.SearchUsers,
	},
	{
		"/users/list",
		http.MethodGet,
		controllers.ListUsers,
	},
	{
		"/forgot",
		http.MethodGet,
		controllers.ForgotUser,
	},
	{
		"/recover",
		http.MethodGet,
		controllers.GetRecoverUser,
	},
	{
		"/recover",
		http.MethodPost,
		controllers.RecoverPasswordUser,
	},
	{
		"/verify",
		http.MethodGet,
		controllers.VerifyEmailUser,
	},
}

var login = []Route{
	{
		"/login",
		http.MethodPost,
		controllers.Login,
	},
	{
		"/admins/login",
		http.MethodPost,
		controllers.AdminLogin,
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
		"/games/similar",
		http.MethodGet,
		controllers.GetSimilar,
	},
	{
		"/games",
		http.MethodPost,
		controllers.CreateGame,
	},
	{
		"/games",
		http.MethodPut,
		controllers.UpdateGame,
	},
	{
		"/games",
		http.MethodDelete,
		controllers.DeleteGame,
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
		"/follow",
		http.MethodDelete,
		controllers.Unfollow,
	},
	{
		"/follow",
		http.MethodGet,
		controllers.GetFollow,
	},
}

var mail = []Route{
	{
		"/mail",
		http.MethodPost,
		controllers.SendMail,
	},
}

var admin = []Route{
	{
		"/admins",
		http.MethodGet,
		controllers.GetAdmin,
	},
	{
		"/admins",
		http.MethodPost,
		controllers.CreateAdmin,
	},
	{
		"/admins",
		http.MethodPatch,
		controllers.EditAdmin,
	},
	{
		"/admins",
		http.MethodPut,
		controllers.EditAdminPassword,
	},
	{
		"/admins",
		http.MethodDelete,
		controllers.DeleteAdmin,
	},
	{
		"/admins/search",
		http.MethodGet,
		controllers.SearchAdmins,
	},
	{
		"/admins/list",
		http.MethodGet,
		controllers.ListAdmins,
	},
	{
		"/admins/forgot",
		http.MethodGet,
		controllers.ForgotAdmin,
	},
	{
		"/admins/recover",
		http.MethodGet,
		controllers.GetRecoverAdmin,
	},
	{
		"/admins/recover",
		http.MethodPost,
		controllers.RecoverPasswordAdmin,
	},
}

var recommended = []Route{
	{
		"/recommendeds",
		http.MethodGet,
		controllers.GetRecommended,
	},
	{
		"/recommendeds/search",
		http.MethodGet,
		controllers.SearchRecommended,
	},
	{
		"/recommendeds",
		http.MethodPost,
		controllers.CreateRecommended,
	},
	{
		"/recommendeds",
		http.MethodPut,
		controllers.ApproveRecommended,
	},
	{
		"/recommendeds",
		http.MethodDelete,
		controllers.DeleteRecommended,
	},
}

var favorite = []Route{
	{
		"/favorite",
		http.MethodPost,
		controllers.FavoriteGame,
	},
	{
		"/favorite",
		http.MethodDelete,
		controllers.UnFavoriteGame,
	},
	{
		"/favorite/user",
		http.MethodGet,
		controllers.GetFavoritesByUser,
	},
	{
		"/favorite/game",
		http.MethodGet,
		controllers.GetFavoritesByGame,
	},
}

var image = []Route{
	{
		"/image/:id",
		http.MethodGet,
		controllers.GetImage,
	},
}