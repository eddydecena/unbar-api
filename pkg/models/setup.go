package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("unbar.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to create table")
	}

	DB = database
}
