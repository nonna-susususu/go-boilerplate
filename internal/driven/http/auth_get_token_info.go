package http

import (
	"fmt"

	"github.com/fastworkco/go-boilerplate/internal/domain"
	"github.com/go-resty/resty/v2"
)

func (a *AuthClient) GetTokenInfo(token string) (domain.AuthTokenData, error) {
	var response AuthTokenResponse

	client := resty.New()
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	r, err := client.R().
		SetHeaders(headers).
		SetResult(&response).
		Post(fmt.Sprintf("%s/auth/v2/token.validate", a.config.Endpoint))
	if err != nil {
		return domain.AuthTokenData{}, err
	}

	if r.IsError() {
		return domain.AuthTokenData{}, nil
	}

	return response.Data.ToDomain(), nil
}
