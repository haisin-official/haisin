package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ajax() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Header
		// X-Requested-With: XMLHttpRequest
		XRequestedWith := c.GetHeader("X-Requested-With")
		if XRequestedWith != "XMLHttpRequest" {
			c.AbortWithStatus(http.StatusNotFound)
			c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
			return
		}

		c.Next()
	}
}
