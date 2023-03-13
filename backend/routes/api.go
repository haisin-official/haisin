package routes

import (
	"github.com/gin-gonic/gin"
	GaController "github.com/haisin-official/haisin/app/controllers/Ga"
	OAuthController "github.com/haisin-official/haisin/app/controllers/OAuth"
	SlugController "github.com/haisin-official/haisin/app/controllers/Slug"
	UserController "github.com/haisin-official/haisin/app/controllers/User"

	"github.com/haisin-official/haisin/app/middleware"
)

func Router(router *gin.Engine) {
	v1 := router.Group("api/v1")
	v1.Use(middleware.Ajax())

	oauth := v1.Group("/oauth/google")
	{
		controller := OAuthController.OAuthController{}
		oauth.GET("/redirect", controller.Redirect)
		oauth.POST("/callback", controller.Callback)
	}

	user := v1.Group("/user")
	user.Use(middleware.Authenticate())
	{
		controller := UserController.UserController{}
		user.GET("/me", controller.UserMeGet)
		user.POST("/logout", controller.UserLogoutPost)
	}

	slug := v1.Group("/slug")
	slug.Use(middleware.Authenticate())
	{
		controller := SlugController.SlugController{}
		slug.GET("/", controller.SlugGet)
		slug.POST("/", controller.SlugPost)
	}

	ga := v1.Group("/ga")
	ga.Use(middleware.Authenticate())
	{
		controller := GaController.GaController{}
		ga.GET("/", controller.GetGa)
		ga.POST("/", controller.PostGa)
		ga.DELETE("/", controller.DeleteGa)
	}
}
