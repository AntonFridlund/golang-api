package main

import (
	"main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func main() {
	app := fiber.New()

	app.Use(helmet.New()) // Use helmet.

	app.Use(logger.New())

	routes.Setup(app) // Setup routes.

	app.Listen(":8080")
}
