package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

// GET /books
// Get all books
func FindBooks(context *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	context.JSON(http.StatusOK, gin.H{"data": books})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// POST /books
// Create new book
func CreateBook(context *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	context.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books/:id
// Find a book
func FindBook(context *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// PATCH /books/:id
// Update a book
func UpdateBook(context *gin.Context) {
	// Get model if it exists
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(context *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("The book %s (%v) has been deleted", book.Title, book.ID)})
}
