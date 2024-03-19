package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/kudoabhijeet/go-rest/internal/app/models"
)

func connect(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	
	db.AutoMigrate(&Todo{})
}