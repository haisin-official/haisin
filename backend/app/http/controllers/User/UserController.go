package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type UserController struct{}

func (UserController) GetUserMe(c *gin.Context) {

	ckey := config.GetEnv("SESSION_KEY")
	data, err := session.GetSession(c, ckey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.SessionId, data.UserId)
	// c.JSON(http.StatusOK, userId)
}
