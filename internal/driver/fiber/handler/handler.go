package handler

import (
	"github.com/fastworkco/go-boilerplate/internal/driver/fiber/handler/health"
	"github.com/fastworkco/go-boilerplate/internal/driver/fiber/handler/todo"
)

type HandlersDependencies struct {
	HealthHandler *health.HealthHandler
	TodoHandler   *todo.TodoHandler
}

type Handlers struct {
	HealthHandler *health.HealthHandler
	TodoHandler   *todo.TodoHandler
}

func NewHandlers(handlersDependencies HandlersDependencies) *Handlers {
	return &Handlers{
		HealthHandler: handlersDependencies.HealthHandler,
	}
}
