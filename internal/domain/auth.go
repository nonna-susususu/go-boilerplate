package domain

type AuthTokenData struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
