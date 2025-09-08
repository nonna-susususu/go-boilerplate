package postgres

import (
	"github.com/fastworkco/go-boilerplate/internal/domain"
	"gorm.io/gorm"
)

type Todo struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func (t *Todo) ToDomain() domain.Todo {
	return domain.Todo{
		Task:   t.Task,
		IsDone: t.IsDone,
	}
}

type TodoRepository struct {
	baseGormRepository
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}
