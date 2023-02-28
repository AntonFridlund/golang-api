package v1

import (
	"github.com/gofiber/fiber/v2"
)

// Api router.
func Router(router fiber.Router) {
	router.Route("/users", userRouter)
}
