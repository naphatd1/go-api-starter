package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Set up routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server on port 4000
	if err := app.Listen(":4000"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
