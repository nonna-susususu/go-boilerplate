package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// HealthResponse represents the health check response structure
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// HealthCheck handles the health check endpoint
func (h Handler) HealthCheck(c *fiber.Ctx) error {
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
	}
	return c.JSON(response)
}
