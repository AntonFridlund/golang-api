package user

import "github.com/gofiber/fiber/v2"

func NewUser(c *fiber.Ctx) error {
	return c.SendString("YOU MADE IT! We are now at: \"/user/new\" 🎉")
}
