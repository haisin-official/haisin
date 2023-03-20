package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Slug"
	usecases "github.com/haisin-official/haisin/app/usecases/Slug"
	"github.com/haisin-official/haisin/config/session"
)

type SlugController struct{}

func (SlugController) SlugGet(c *gin.Context) {
	cKey := os.Getenv("SESSION_KEY")
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

	result, code, err := usecases.SlugUseCase{}.SlugGetAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

func (SlugController) SlugPost(c *gin.Context) {
	cKey := os.Getenv("SESSION_KEY")
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
	// SLUGのValidationを行う
	reqBody := new(requests.SlugPostRequestBody)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	req := requests.SlugPostRequest{
		UserID: u,
		Slug:   reqBody.Slug,
	}

	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.SlugUseCase{}.SlugPostAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}
