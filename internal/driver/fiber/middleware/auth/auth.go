package auth

import (
	"errors"

	"github.com/fastworkco/go-boilerplate/internal/domain"
	"go.uber.org/zap"
)

const (
	ContextRequestToken string = "RequestToken"
	ContextAuthUserID   string = "AuthUserID"
	ContextAuthRole     string = "AuthRole"
)

var (
	ErrorIdentityNotFound  error = errors.New("user identity not found or incomplete in token")
	ErrorAuthTokenNotFound error = errors.New("auth token not found in context")
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
