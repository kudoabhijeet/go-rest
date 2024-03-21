package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"rest/internal/app/models"
)

var DB *gorm.DB
func Connect(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic("Failed to migrate database!")
	}
	DB = db
}