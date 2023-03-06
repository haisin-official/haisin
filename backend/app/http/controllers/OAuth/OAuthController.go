package contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/haisin-official/haisin/app/usecases/OAuth"
)

type OAuthController struct{}

func (OAuthController) Redirect(c *gin.Context) {
	result, code, err := usecases.OAuthUseCases{}.RedirectAction()

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	} else {
		c.JSON(code, result)
	}
}
