package user

import "github.com/gofiber/fiber/v2"

func User(c *fiber.Ctx) error {
	return c.SendString("You are now at: \"/user\"\nHere you could display some user information.")
}

func NewUser(c *fiber.Ctx) error {
	return c.SendString("You are now at: \"/user/new\".\nHere you could create a new user.")
}
