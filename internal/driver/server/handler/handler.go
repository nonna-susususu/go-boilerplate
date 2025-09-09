package handler

import (
	"errors"

	authMiddleware "github.com/fastworkco/go-boilerplate/internal/driver/server/middleware/auth"
	app "github.com/fastworkco/go-boilerplate/internal/service"
	"github.com/fastworkco/go-boilerplate/internal/service/todo"
	"github.com/gofiber/fiber/v2"
)

// APICommonError represents the common error response
type APICommonError struct {
	Error  string `json:"code"`
	Detail string `json:"detail"`
}

type HandlerDependencies struct {
	TodoService todo.TodoService
}

type Handler struct {
	TodoService todo.TodoService
}

func NewHandler(handlerDependencies HandlerDependencies) *Handler {
	return &Handler{
		TodoService: handlerDependencies.TodoService,
	}
}

func (h *Handler) GetUserAndToken(c *fiber.Ctx) (string, string, error) {
	var userID string
	var tokenString string

	userID, ok := c.Locals(authMiddleware.ContextAuthUserID).(string)
	if !ok {
		return "", "", c.Status(fiber.StatusUnauthorized).JSON(APICommonError{
			Error:  "Unauthorized",
			Detail: "User identity not found or incomplete in token",
		})
	}

	tokenString, ok = c.Locals(authMiddleware.ContextRequestToken).(string)
	if !ok || tokenString == "" {
		return "", "", c.Status(fiber.StatusUnauthorized).JSON(APICommonError{
			Error:  "Unauthorized",
			Detail: "Auth token not found in context",
		})
	}

	return userID, tokenString, nil
}

func ServiceErrorToHTTPResponse(e error) (int, APICommonError) {
	if e == nil {
		return fiber.StatusOK, APICommonError{}
	}

	switch {
	// service error
	case errors.Is(e, app.ErrUnauthorized):
		return fiber.StatusUnauthorized, APICommonError{
			Error:  "UNAUTHORIZED",
			Detail: "Unauthorized access, please check your credentials",
		}
	// otherwise
	default:
		return fiber.StatusInternalServerError, APICommonError{
			Error:  "INTERNAL_SERVER_ERROR",
			Detail: "Internal server error",
		}
	}
}
