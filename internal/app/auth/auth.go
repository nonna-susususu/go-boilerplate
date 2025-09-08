package auth

import (
	"github.com/fastworkco/go-boilerplate/internal/domain"
)

type AuthClient interface {
	GetTokenInfo(token string) (domain.AuthTokenData, error)
}

type AuthServiceDependencies struct {
	AuthClient AuthClient
}

type AuthService interface {
	GetTokenInfo(token string) (domain.AuthTokenData, error)
}

type authService struct {
	authClient AuthClient
}

func NewAuthService(authServiceDependencies AuthServiceDependencies) AuthService {
	return &authService{
		authClient: authServiceDependencies.AuthClient,
	}
}
