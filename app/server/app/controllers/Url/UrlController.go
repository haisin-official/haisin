package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	requests "github.com/haisin-official/haisin/app/requests/Url"
	usecases "github.com/haisin-official/haisin/app/usecases/Url"
	"github.com/haisin-official/haisin/ent/service"

	"github.com/go-playground/validator/v10"
)

type UrlController struct{}

func (UrlController) UrlSlugGet(c *gin.Context) {
	slug := c.Param("slug")
	// slugが要件を満たしているのか確認する
	if err := validator.New().Var(slug, "required,alphanum,min=4,max=30"); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	req := requests.UrlSlugGetRequest{
		Slug: slug,
	}
	result, code, err := usecases.UrlUseCase{}.UrlSlugGetAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}

func (UrlController) UrlSlugServiceGet(c *gin.Context) {
	slug := c.Param("slug")
	name := c.Param("service")
	// slugが要件を満たしているのか確認する
	if err := validator.New().Var(slug, "required,alphanum,min=4,max=30"); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	// nameが要件を満たしているのか確認する
	if err := service.ServiceValidator(service.Service(name)); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
		return
	}

	req := requests.UrlSlugServiceGetRequest{
		Slug:    slug,
		Service: name,
	}
	result, code, err := usecases.UrlUseCase{}.UrlSlugServiceGetAction(req)

	if err != nil {
		c.AbortWithStatus(code)
		c.JSON(code, gin.H{"error": http.StatusText(code)})
		return
	}

	c.JSON(code, result)
}
