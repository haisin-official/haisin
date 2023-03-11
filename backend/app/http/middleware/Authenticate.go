package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// セッションが存在するか確認
		cKey := config.GetEnv("SESSION_KEY")
		data, httpCode, err := session.GetSession(c, cKey)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(httpCode)
			c.JSON(httpCode, gin.H{"error": http.StatusText(httpCode)})
			return
		}
		// セッションの値が空の場合401エラーを返す
		if data.SessionId == "" || data.UserId == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		}
		c.Next()
	}
}
