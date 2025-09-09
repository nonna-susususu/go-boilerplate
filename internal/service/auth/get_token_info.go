package auth

import (
	"github.com/fastworkco/go-boilerplate/internal/domain"
)

func (as *authService) GetTokenInfo(token string) (domain.AuthTokenData, error) {
	return as.authClient.GetTokenInfo(token)
}
