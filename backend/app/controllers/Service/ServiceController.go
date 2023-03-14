package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Service"
	usecases "github.com/haisin-official/haisin/app/usecases/Service"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type ServiceController struct{}

func (ServiceController) ServiceGet(c *gin.Context) {
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

	req := requests.ServiceGetRequest{
		UserID: u,
	}

	result, code, err := usecases.ServiceUseCase{}.ServiceGetAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

func (ServiceController) ServicePost(c *gin.Context) {
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

	reqBody := new(requests.ServicePostRequestBody)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	req := requests.ServicePostRequest{
		UserID:  u,
		Service: reqBody.Service,
		Url:     reqBody.Url,
	}

	result, code, err := usecases.ServiceUseCase{}.ServicePostAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}
