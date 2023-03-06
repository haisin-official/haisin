package routes

import (
	"github.com/gin-gonic/gin"
	OAuthController "github.com/haisin-official/haisin/app/http/controllers/OAuth"
)

func Router(router *gin.Engine) {
	v1 := router.Group("api/v1")

	oauth := v1.Group("/oauth/google")
	{
		controller := OAuthController.OAuthController{}
		oauth.GET("/redirect", controller.Redirect)
	}
}
