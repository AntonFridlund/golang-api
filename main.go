package main

import (
	"errors"
	"log"
	"main/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func main() {
	// Server configuration.
	var config = fiber.Config{
		ErrorHandler: errorHandler,
	}
	app := fiber.New(config)
	app.Use(logger.New()) // Use logger.
	app.Use(helmet.New()) // Use helmet.

	// Setup routes.
	routes.Setup(app)

	// Start server.
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}

// Custom error handler.
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError // Default error code.

	// Replace default error code if a specific code exists.
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Supress dangerous messages.
	var message string
	switch code {
	case 500:
		message = fiber.ErrInternalServerError.Message
	case 404:
		message = fiber.ErrNotFound.Message
	default:
		message = err.Error()
	}

	// Set Content-Type: text/plain; charset=utf-8.
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message.
	return c.Status(code).SendString(message)
}
