package v1

import (
	controller "main/controllers/user"

	"github.com/gofiber/fiber/v2"
)

// User router.
func userRouter(router fiber.Router) {
	router.Get("/new", controller.NewUser)
	router.Get("/", controller.User)
}
