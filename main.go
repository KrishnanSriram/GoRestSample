package main

import (
	approuter "SimpleRestAPI/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	approuter.AddBookRoutes(router)
	approuter.AddHealthRoutes(router)
	router.Use(approuter.ErrorHandler)

	router.Run(":8080")
}
