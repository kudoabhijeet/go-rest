package main

import (
	"log"
	"rest/internal/app/database"
	"rest/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Get("/todos", handlers.GetTodos)
	app.Post("/todos", handlers.CreateTodo)
	

	log.Fatal(app.Listen(":3000"))
}