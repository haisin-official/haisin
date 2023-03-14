package routes

import (
	"github.com/gin-gonic/gin"
	GaController "github.com/haisin-official/haisin/app/controllers/Ga"
	OAuthController "github.com/haisin-official/haisin/app/controllers/OAuth"
	ServiceController "github.com/haisin-official/haisin/app/controllers/Service"
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
		oauth.GET("/redirect", controller.RedirectGet)
		oauth.POST("/callback", controller.CallbackPost)
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
		ga.GET("/", controller.GaGet)
		ga.POST("/", controller.GaPost)
		ga.DELETE("/", controller.GaDelete)
	}

	service := v1.Group("/service")
	service.Use(middleware.Authenticate())
	{
		controller := ServiceController.ServiceController{}
		service.GET("/", controller.ServiceGet)
		service.POST("/", controller.ServicePost)
		service.DELETE("/:name", controller.ServiceDelete)
	}
}
