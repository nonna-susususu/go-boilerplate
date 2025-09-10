package auth

import "github.com/fastworkco/go-boilerplate/internal/domain"

type AuthTokenData struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func (a AuthTokenData) ToDomain() domain.AuthTokenData {
	return domain.AuthTokenData(a)
}

type AuthTokenResponse struct {
	Data AuthTokenData `json:"data"`
}

type AuthClientConfig struct {
	Endpoint string
}

type AuthClient struct {
	config AuthClientConfig
}

func NewAuthClient(config AuthClientConfig) *AuthClient {
	return &AuthClient{
		config: config,
	}
}
