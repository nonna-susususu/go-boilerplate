package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"

	"github.com/fastworkco/common-go/load_config/v1"
	"github.com/fastworkco/common-go/log/v1"
	"github.com/fastworkco/common-go/telemetry/v1"
	"github.com/fastworkco/go-boilerplate/internal/config"
	authClient "github.com/fastworkco/go-boilerplate/internal/driven/http/auth"
	"github.com/fastworkco/go-boilerplate/internal/driven/postgres"
	"github.com/fastworkco/go-boilerplate/internal/driver/fiber"
	"github.com/fastworkco/go-boilerplate/internal/driver/fiber/handler"
	"github.com/fastworkco/go-boilerplate/internal/driver/fiber/handler/health"
	"github.com/fastworkco/go-boilerplate/internal/metrics"

	todoRepo "github.com/fastworkco/go-boilerplate/internal/driven/postgres/todo"
	todoHandler "github.com/fastworkco/go-boilerplate/internal/driver/fiber/handler/todo"
	authService "github.com/fastworkco/go-boilerplate/internal/service/auth"
	"github.com/fastworkco/go-boilerplate/internal/service/todo"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	cfg := load_config.New[config.Config]()
	log.Init(cfg.AppConfig.Env, cfg.AppConfig.AppName)

	closeTelemetry, err := telemetry.Initialize(ctx, cfg.Telemetry)
	if err != nil {
		log.Logger.Fatal("[main]: unable to initialize telemetry", zap.Error(err))
	}
	defer closeTelemetry()

	if err := metrics.InitializeMetrics(cfg.AppConfig.AppName); err != nil {
		log.Logger.Fatal("[main]: unable to initialize metrics", zap.Error(err))
	}

	db := postgres.InitGormPostgres(cfg.AppConfig.Env, cfg.DatabaseConfig, log.Logger)
	todoRepository := todoRepo.NewTodoRepository(db)

	todoService := todo.NewTodoService(todo.TodoServiceDependencies{
		TodoRepository: todoRepository,
	})

	healthHandler := health.NewHealthHandler()
	_todoHandler := todoHandler.NewTodoHandler(todoHandler.TodoHandlerDependencies{
		TodoService: todoService,
	})

	// setup handlers
	handler := handler.NewHandlers(handler.HandlersDependencies{
		HealthHandler: healthHandler,
		TodoHandler:   _todoHandler,
	})

	authClient := authClient.NewAuthClient(authClient.AuthClientConfig{
		Endpoint: cfg.RemoteConfig.FastworkAuthProviderURL,
	})

	authService := authService.NewAuthService(authService.AuthServiceDependencies{
		AuthClient: authClient,
	})

	srv := fiber.New(cfg.AppConfig, authService, handler, log.Logger)

	wg := sync.WaitGroup{}

	wg.Go(func() {
		if err := srv.Start(ctx); err != nil {
			log.Logger.Error("[main]: server error", zap.Error(err))
		}
	})

	<-ctx.Done()
	log.Logger.Info("[main]: shutting down server...")
	wg.Wait()
	log.Logger.Info("[main]: consumers and server shut down successfully")
}
