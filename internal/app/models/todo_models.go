package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID int `json:"id"`
	Todo string `json:"todo"`
}

