package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/http/requests/User"
	usecases "github.com/haisin-official/haisin/app/usecases/User"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type UserController struct{}

// ログイン中のユーザーデータを取得
func (UserController) GetUserMe(c *gin.Context) {

	cKey := config.GetEnv("SESSION_KEY")
	data, httpCode, err := session.GetSession(c, cKey)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(httpCode)
		c.JSON(httpCode, gin.H{"error": http.StatusText(httpCode)})
		return
	}
	// セッションから取得したユーザーIDをUUIDに変換する
	u, err := uuid.Parse(data.UserId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	// リクエスト用のデータを構築
	req := requests.UserGetRequest{
		UserID: u,
	}

	// dataからUserIDを取得
	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.UserUseCases{}.GetUserAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

// ログアウトする
func (UserController) Logout(c *gin.Context) {
	cKey := config.GetEnv("SESSION_KEY")
	if err := session.DeleteSession(c, cKey); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
