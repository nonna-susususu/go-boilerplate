package postgres

import (
	"fmt"

	"github.com/fastworkco/go-boilerplate/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGormPostgres(env string, c config.DatabaseConfig, log *zap.Logger) *gorm.DB {
	sslmode := "disable"
	if env != "local" {
		sslmode = "require"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Bangkok", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, sslmode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to postgres", zap.Error(err))
	}

	return db
}
