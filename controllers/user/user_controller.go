package user

import "github.com/gofiber/fiber/v2"

func User(c *fiber.Ctx) error {
	return c.SendString("You are now at: \"/users\".")
}

func NewUser(c *fiber.Ctx) error {
	return c.SendString("You are now at: \"/users/new\".")
}
