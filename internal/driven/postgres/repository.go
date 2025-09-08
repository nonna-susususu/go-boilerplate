package postgres

import (
	"context"

	"gorm.io/gorm"
)

type ContextKey string

var GormTransactionContextKey ContextKey = ContextKey("gorm_tx")

type GormTransactionControl struct {
	db *gorm.DB
}

func NewGormTransactionControl(db *gorm.DB) *GormTransactionControl {
	return &GormTransactionControl{db: db}
}

// Do executes the given function within a GORM transaction.
func (u *GormTransactionControl) Do(ctx context.Context, fn func(txCtx context.Context) error) error {
	if _, ok := ctx.Value(GormTransactionContextKey).(*gorm.DB); ok {
		return fn(ctx)
	}

	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txCtx := context.WithValue(ctx, GormTransactionContextKey, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

type baseGormRepository struct{}

func (r *baseGormRepository) getDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(GormTransactionContextKey).(*gorm.DB); ok {
		return tx
	}

	return db
}
