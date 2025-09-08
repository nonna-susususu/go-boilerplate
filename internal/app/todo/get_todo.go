package todo

import (
	"context"

	"github.com/fastworkco/go-boilerplate/internal/domain"
)

func (s *todoService) GetAllTodo(ctx context.Context) ([]domain.Todo, error) {
	return s.todoRepository.GetAll(ctx)
}
