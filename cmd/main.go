package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/kudoabhijeet/go-rest/internal/app/handlers
	
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("/new", handlers.createTodo)
	app.Get("/:id", handlers.getTodo)
	app.Put("/:id", handlers.updateTodo)
	app.Delete("/:id", handlers.deleteTodo)
	

	log.Fatal(app.Listen(":3000"))
}