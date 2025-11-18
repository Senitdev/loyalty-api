package dto

type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	Admin     bool   `json:"admin"`
}
