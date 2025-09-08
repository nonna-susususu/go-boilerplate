package todo

import (
	"context"

	"github.com/fastworkco/go-boilerplate/internal/domain"
)

type TodoRepository interface {
	GetAll(ctx context.Context) ([]domain.Todo, error)
}

type TodoServiceDependencies struct {
	TodoRepository TodoRepository
}

type TodoService interface {
	GetAllTodo(ctx context.Context) ([]domain.Todo, error)
}

type todoService struct {
	todoRepository TodoRepository
}

func NewTodoService(todoDependencies TodoServiceDependencies) TodoService {
	return &todoService{
		todoRepository: todoDependencies.TodoRepository,
	}
}
