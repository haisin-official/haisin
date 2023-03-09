package contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/haisin-official/haisin/app/usecases/OAuth"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type OAuthController struct{}

func (OAuthController) Redirect(c *gin.Context) {
	result, code, err := usecases.OAuthUseCases{}.RedirectAction()

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}
	c.JSON(code, result)
}

func (OAuthController) Callback(c *gin.Context) {
	state := c.Query("state")
	OAuthCode := c.Query("code")
	result, code, err := usecases.OAuthUseCases{}.CallbackAction(state, OAuthCode)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": err})
		return
	}

	// sessionにユーザーIDを保管
	ckey := config.GetEnv("SESSION_SECERT_ENCRYPTION_KEY")
	data := new(session.Store)
	data.SessionId = config.GetEnv("SESSION_KEY")
	data.UserId = result.User.Uuid.String()

	session.NewSession(c, ckey, *data)

	c.JSON(code, result)
}
