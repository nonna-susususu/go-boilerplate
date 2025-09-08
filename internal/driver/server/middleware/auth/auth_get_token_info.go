package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetTokenInfo middleware to extract token info from JWT
func (am AuthMiddleware) GetTokenInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(ContextAuthUserID, "")
		c.Locals(ContextAuthRole, "")
		c.Locals(ContextRequestToken, "")

		authorization := c.Get("Authorization")
		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error_code": "UNAUTHORIZED",
				"message":    "Authorization header required",
			})
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		if token == authorization || token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error_code": "UNAUTHORIZED",
				"message":    "Malformed token or missing token",
			})
		}

		data, err := am.authProvider.GetTokenInfo(token)
		if err != nil {
			return c.Next()
		}

		c.Locals(ContextAuthUserID, data.UserID)
		c.Locals(ContextAuthRole, data.Role)
		c.Locals(ContextRequestToken, token)

		return c.Next()
	}
}
