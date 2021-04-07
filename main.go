package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nickvirden/go-gin-books/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}
