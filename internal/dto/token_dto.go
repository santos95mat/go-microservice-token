package dto

type CreateTokenDTO struct {
	UserID string `json:"user_id"`
}

type OutputTokenDTO struct {
	Token string `json:"token"`
}

type ValidateTokenDTO struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
