package todo

import (
	"context"
	"errors"

	"github.com/fastworkco/go-boilerplate/internal/domain"
	"github.com/fastworkco/go-boilerplate/internal/driven/postgres"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (r *TodoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	var todos []*Todo

	db := postgres.GetDBFromCtx(ctx, r.db)

	if err := db.Find(&todos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	result := lo.Map(todos, func(t *Todo, _ int) domain.Todo {
		return t.ToDomain()
	})

	return result, nil
}
