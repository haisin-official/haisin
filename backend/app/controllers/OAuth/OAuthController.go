package contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requests "github.com/haisin-official/haisin/app/requests/OAuth"
	usecases "github.com/haisin-official/haisin/app/usecases/OAuth"
	"github.com/haisin-official/haisin/app/utils"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type OAuthController struct{}

func (OAuthController) RedirectGet(c *gin.Context) {
	req := requests.RedirectGetRequest{}
	result, code, err := usecases.OAuthUseCase{}.RedirectGetAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}
	c.JSON(code, result)
}

func (OAuthController) CallbackPost(c *gin.Context) {
	req := requests.CallbackPostRequest{
		State: c.Query("state"),
		Code:  c.Query("code"),
	}
	result, code, err := usecases.OAuthUseCase{}.CallbackPostAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": err})
		return
	}

	// 一度セッションを解除する
	cKey := config.GetEnv("SESSION_KEY")
	session.DeleteSession(c, cKey)

	// sessionにユーザーIDを保管
	cKey = config.GetEnv("SESSION_KEY")
	data := session.Store{
		SessionId: utils.GenRandToken(),
		UserId:    result.User.Uuid.String(),
	}
	session.NewSession(c, cKey, data)

	c.JSON(code, result)
}
