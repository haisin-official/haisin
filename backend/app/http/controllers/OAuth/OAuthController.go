package contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requests "github.com/haisin-official/haisin/app/http/requests/OAuth"
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
	req := requests.CallbackRequest{
		State: c.Query("state"),
		Code:  c.Query("code"),
	}
	result, code, err := usecases.OAuthUseCases{}.CallbackAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": err})
		return
	}

	// 一度セッションを解除する
	cKey := config.GetEnv("SESSION_SECRET_ENCRYPTION_KEY")
	session.DeleteSession(c, cKey)

	// sessionにユーザーIDを保管
	ckey := config.GetEnv("SESSION_SECERT_ENCRYPTION_KEY")
	data := new(session.Store)
	data.SessionId = config.GetEnv("SESSION_KEY")
	data.UserId = result.User.Uuid.String()

	session.NewSession(c, ckey, *data)

	c.JSON(code, result)
}
