package routes

import (
	controller "main/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func User() *fiber.App {
	router := fiber.New()                  // Register a new router.
	router.Get("/", controller.User)       // Register a new route for GET in /user.
	router.Get("/new", controller.NewUser) // Register a new route for GET in /user.
	return router                          // Return the router in order to mount it.
}
