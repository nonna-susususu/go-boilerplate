package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/fastworkco/common-go/log/v1"
	"github.com/fastworkco/go-boilerplate/internal/config"
	"github.com/fastworkco/go-boilerplate/internal/driver/server/handler"
	"github.com/fastworkco/go-boilerplate/internal/driver/server/middleware"
	authMiddleware "github.com/fastworkco/go-boilerplate/internal/driver/server/middleware/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

// Server represents the HTTP server
type Server struct {
	app     *fiber.App
	appName string
	port    int
}

// New creates a new server instance
func New(
	appConfig config.AppConfig,
	authProvider authMiddleware.AuthProvider,
	handler *handler.Handler,
	logger *zap.Logger,
) *Server {
	app := fiber.New(fiber.Config{
		AppName:      appConfig.AppName,
		ErrorHandler: log.CreateFiberLogContextErrorHandler(),
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(appConfig.Cors, ","),
		AllowHeaders: strings.Join(appConfig.CorsHeader, ","),
	}))
	app.Use(log.CreateFiberLogContextAttachment())
	app.Use(log.CreateFiberRequestLogMiddleware(logger))
	app.Use(recover.New())

	// Setup routes
	app.Get("/health", handler.HealthCheck)

	// api group v1 (external)
	authMiddleware := authMiddleware.NewAuthMiddleware(authProvider, logger)
	v1 := app.Group("/api/v1")
	v1.Use(authMiddleware.GetTokenInfo())
	v1.Use(authMiddleware.AuthUserGuard())

	v1.Get("/todo.getAll", handler.ListTodo)

	app.Use(middleware.PageNotFound)

	return &Server{
		app:     app,
		appName: appConfig.AppName,
		port:    appConfig.Port,
	}
}

// Start starts the server and handles graceful shutdown
func (s *Server) Start(ctx context.Context) error {
	// Listen for context done for graceful shutdown
	shutdownComplete := make(chan struct{})
	go func() {
		<-ctx.Done()
		log.Logger.Info("[Service.Start] Shutting down server...")

		_ = s.app.Shutdown()
		shutdownComplete <- struct{}{}
	}()

	// Start the server
	err := s.app.Listen(fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Logger.Error("[Service.Start] Failed to start server", zap.Error(err))
		return fmt.Errorf("failed to start server: %w", err)
	}

	<-shutdownComplete
	log.Logger.Info("[Service.Start] Server shut down gracefully")

	return nil
}
