package handlers

import (
	"rest/internal/app/services"
	"rest/internal/app/models"
	"github.com/gofiber/fiber/v2"
)


func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	if err := services.CreateTodoService(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating todo",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}


func GetTodos(c *fiber.Ctx) error {
	todos, err := services.GetTodoService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching todos",
		})
	}
	return c.JSON(todos)
}
// func deleteTodo(c *fiber.Ctx) error {
// }

// func updateTodo(c *fiber.Ctx) error {
	
// }