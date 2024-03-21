package services

import (
	"rest/internal/app/database"
	"rest/internal/app/models"
)

func GetTodoService() ([]models.Todo, error) {
	var todos []models.Todo
	if err:= database.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
func CreateTodoService(todo models.Todo) error {
	if err:= database.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}


