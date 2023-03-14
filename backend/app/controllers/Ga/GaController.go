package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Ga"
	usecases "github.com/haisin-official/haisin/app/usecases/Ga"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/config/session"
)

type GaController struct{}

func (GaController) GaGet(c *gin.Context) {
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

	req := requests.GaGetRequest{
		UserID: u,
	}

	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.GaUseCase{}.GaGetAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

func (GaController) GaPost(c *gin.Context) {
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
	// ValidationとBindを実行
	reqBody := new(requests.GaPostRequestBody)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	req := requests.GaPostRequest{
		UserID: u,
		Ga:     reqBody.Ga,
	}

	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.GaUseCase{}.GaPostAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

func (GaController) GaDelete(c *gin.Context) {
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

	req := requests.GaDeleteRequest{
		UserID: u,
	}

	fmt.Println(data.SessionId, data.UserId)

	result, code, err := usecases.GaUseCase{}.GaDeleteAction(req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(code, result)

}
