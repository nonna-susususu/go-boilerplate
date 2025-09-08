package handler

import (
	"fmt"

	"github.com/fastworkco/common-go/log/v1"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) ListTodo(c *fiber.Ctx) error {
	ctx := c.UserContext()
	lc := log.GetLogContext(ctx)

	// get user id and token from context
	userID, _, err := h.GetUserAndToken(c)
	lc.Logger().Info(fmt.Sprintf("[handler.listTodo]: request by userID: %s", userID))
	if err != nil {
		lc.Logger().Error("[handler.listTodo]: Failed to get user ID and token from context",
			zap.Error(err),
		)
		return c.Status(fiber.StatusUnauthorized).JSON(APICommonError{
			Error:  "UNAUTHORIZED",
			Detail: "User identity not found or incomplete in token",
		})
	}

	result, err := h.TodoService.GetAllTodo(ctx)
	if err != nil {
		code, err := ServiceErrorToHTTPResponse(err)
		return c.Status(code).JSON(err)
	}

	return c.JSON(result)
}
