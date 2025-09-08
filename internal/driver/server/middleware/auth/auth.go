package auth

import (
	"github.com/fastworkco/go-boilerplate/internal/domain"
	"go.uber.org/zap"
)

const (
	ContextRequestToken string = "RequestToken"
	ContextAuthUserID   string = "AuthUserID"
	ContextAuthRole     string = "AuthRole"
)

type AuthMiddleware struct {
	authProvider AuthProvider
	logger       *zap.Logger
}

type AuthProvider interface {
	GetTokenInfo(token string) (domain.AuthTokenData, error)
}

func NewAuthMiddleware(authProvider AuthProvider, logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		authProvider: authProvider,
		logger:       logger,
	}
}
