package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Create a new Fiber app
	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server on port 3000
	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
