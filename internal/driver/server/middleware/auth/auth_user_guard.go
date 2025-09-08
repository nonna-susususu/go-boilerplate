package auth

import (
	"github.com/gofiber/fiber/v2"
)

// AuthUserGuard middleware to check if the user is authenticated
func (am AuthMiddleware) AuthUserGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authUserID := c.Locals(ContextAuthUserID).(string)

		if authUserID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fiber.Map{
					"code":   "UNAUTHORIZED",
					"detail": "unauthorized",
				},
			})
		}

		return c.Next()
	}
}
