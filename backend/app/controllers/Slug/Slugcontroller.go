package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/http/requests/Slug"
	usecases "github.com/haisin-official/haisin/app/usecases/Slug"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type SlugController struct{}

func (SlugController) GetSlug(c *gin.Context) {
	cKey := config.GetEnv("SESSION_KEY")
	data, httpCode, err := session.GetSession(c, cKey)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(httpCode)
		c.JSON(httpCode, gin.H{"error": http.StatusText(httpCode)})
		return
	}

	u, err := uuid.Parse(data.UserId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	req := requests.SlugGetRequest{
		UserID: u,
	}

	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.SlugUseCases{}.GetSlugAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

func (SlugController) PostSlug(c *gin.Context) {
	cKey := config.GetEnv("SESSION_KEY")
	data, httpCode, err := session.GetSession(c, cKey)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(httpCode)
		c.JSON(httpCode, gin.H{"error": http.StatusText(httpCode)})
		return
	}

	u, err := uuid.Parse(data.UserId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	// リクエストからSLUGが存在するか確認する
	reqBody := new(requests.SlugPostRequestBody)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	req := requests.SlugPostRequest{
		UserID: u,
		Slug:   reqBody.Slug,
	}

	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.SlugUseCases{}.PostSlugAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}
