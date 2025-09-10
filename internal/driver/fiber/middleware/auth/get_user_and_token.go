package auth

import "github.com/gofiber/fiber/v2"

func GetUserAndToken(c *fiber.Ctx) (string, string, error) {
	var userID string
	var tokenString string

	userID, ok := c.Locals(ContextAuthUserID).(string)
	if !ok {
		return "", "", ErrorIdentityNotFound
	}

	tokenString, ok = c.Locals(ContextRequestToken).(string)
	if !ok || tokenString == "" {
		return "", "", ErrorAuthTokenNotFound
	}

	return userID, tokenString, nil
}
