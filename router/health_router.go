package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddHealthRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All is well",
		})
	})
}
