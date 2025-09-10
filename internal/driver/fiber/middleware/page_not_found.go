package middleware

import "github.com/gofiber/fiber/v2"

func PageNotFound(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
