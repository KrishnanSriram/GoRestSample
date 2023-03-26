package router

import (
	"SimpleRestAPI/handler"
	"SimpleRestAPI/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _, bookHandler = handler.CreateBookHandler()

func AddBookRoutes(router *gin.Engine) {
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "list books",
			"data":   bookHandler.List(),
		})
	})

	router.GET("/books/:book_id", func(c *gin.Context) {
		var book_id = c.Param("book_id")
		book, err := bookHandler.Find(book_id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"module": "find book",
				"data":   "No matching book found!!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"module": "find book",
			"data":   book,
		})

	})

	router.POST("/books", func(c *gin.Context) {
		var bookReq model.BookRequest
		if err := c.ShouldBindJSON(&bookReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"module": "Add new book",
				"error":  err.Error(),
			})
			return
		}
		book, err := bookHandler.Add(bookReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"module": "Add new book",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"module": "Add new Book",
			"data":   book,
		})
	})
}
