package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Pages       int64  `json:"pages"`
}

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Pages       int64  `json:"pages" binding:"required"`
}

func getBook(c *gin.Context) {
	var book Book

	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func getBooks(c *gin.Context) {
	var books []Book
	DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func addBook(c *gin.Context) {
	var input Book

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	book := Book{Title: input.Title, Description: input.Description, Author: input.Author, Pages: input.Pages}
	DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func updateBook(c *gin.Context) {
	var book Book
	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func deleteBook(c *gin.Context) {
	var book Book
	if err := DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func main() {
	database, err := gorm.Open(sqlite.Open("unbar.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to create table")
	}

	DB = database

	r := gin.Default()

	r.GET("/books/:id", getBook)
	r.GET("/books", getBooks)
	r.POST("/books", addBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	err = r.Run()

	if err != nil {
		panic("Failed to start server")
	}
}
