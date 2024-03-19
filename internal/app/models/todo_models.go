package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	id int
	todo string
}

