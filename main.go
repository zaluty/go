package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	app := fiber.New()
	todos := []Todo{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env vars")
	}

	PORT := os.Getenv("")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})
	// Create a todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
		// ddd
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"err": "body must not be null"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})
	// Update a to-do
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "erro"})
	})
	//Delete a To-Do
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint((todo.ID)) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(todos)

			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	})
	app.Listen(":" + PORT)
}
