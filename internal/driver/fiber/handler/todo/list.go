package todo

import (
	"fmt"

	"github.com/fastworkco/common-go/log/v1"
	"github.com/fastworkco/go-boilerplate/internal/driver/fiber/handler/common"
	authMiddleware "github.com/fastworkco/go-boilerplate/internal/driver/fiber/middleware/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *TodoHandler) ListTodo(c *fiber.Ctx) error {
	ctx := c.UserContext()
	lc := log.GetLogContext(ctx)

	// get user id and token from context
	userID, _, err := authMiddleware.GetUserAndToken(c)
	lc.Logger().Info(fmt.Sprintf("[handler.listTodo]: request by userID: %s", userID))
	if err != nil {
		lc.Logger().Error("[handler.listTodo]: Failed to get user ID and token from context",
			zap.Error(err),
		)
		return c.Status(fiber.StatusUnauthorized).JSON(common.APICommonError{
			Error:  "UNAUTHORIZED",
			Detail: err.Error(),
		})
	}

	result, err := h.todoService.GetAllTodo(ctx)
	if err != nil {
		code, err := common.ServiceErrorToHTTPResponse(err)
		return c.Status(code).JSON(err)
	}

	return c.JSON(result)
}
