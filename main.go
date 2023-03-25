package main

import (
	approuter "SimpleRestAPI/router"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	approuter.AddBookRoutes(router)
	approuter.AddHealthRoutes(router)
	router.Use(approuter.ErrorHandler)

	router.Run(":8080")
}
