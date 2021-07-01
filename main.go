package main

import (
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Setup(app) // Setup routes.

	app.Listen(":8080")
}
