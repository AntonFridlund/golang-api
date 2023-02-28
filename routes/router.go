package routes

import (
	api "main/routes/api/v1"

	"github.com/gofiber/fiber/v2"
)

// Sets up sub-routes.
func Setup(app *fiber.App) {
	app.Route("/api/v1", api.Router)

	// Default api version.
	app.Route("/api", api.Router)

	// 404 routes.
	app.Use("*", func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
