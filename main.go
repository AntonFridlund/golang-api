package main

import (
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	routes.Setup(app) // Setup routes.

	app.Listen(":8080")
}
