package common

import (
	"errors"

	"github.com/fastworkco/go-boilerplate/internal/service"
	"github.com/gofiber/fiber/v2"
)

// APICommonError represents the common error response
type APICommonError struct {
	Error  string `json:"code"`
	Detail string `json:"detail"`
}

func ServiceErrorToHTTPResponse(e error) (int, APICommonError) {
	if e == nil {
		return fiber.StatusOK, APICommonError{}
	}

	switch {
	// service error
	case errors.Is(e, service.ErrUnauthorized):
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
