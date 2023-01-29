package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/unbar-io/unbar/pkg/handlers"
	"github.com/unbar-io/unbar/pkg/models"
)

func main() {
	// initialize database
	var database *gorm.DB

	dbName, present := os.LookupEnv("DB")
	if present == true && dbName == "postgres" {
		connectionString, present := os.LookupEnv("DB_CONNECTION_STRING")
		if present == false {
			panic("Environment Variable \"DB_CONNECTION_STRING\" not set")
		}

		db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err != nil {
			panic("Failed to connect database")
		}

		database = db
	} else {
		db, err := gorm.Open(sqlite.Open("unbar.db"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect database")
		}
		database = db
	}

	err := database.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Failed to create table")
	}

	models.DB = database

	r := gin.Default()

	r.GET("/books/:id", handlers.FindBook)
	r.GET("/books", handlers.FindBooks)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	err = r.Run()

	if err != nil {
		panic("Failed to start server")
	}
}
