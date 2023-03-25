package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func AddErrorRouter(router *gin.Engine) {
//
//}

func ErrorHandler(c *gin.Context) {
	c.Next()
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "The path you are looking for is not not available. Try again later",
	})
	return
}
