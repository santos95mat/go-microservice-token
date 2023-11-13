package dto

import "time"

type CreateTokenDTO struct {
	UserID     string `json:"user_id"`
	Validation int    `json:"validation"`
}

type OutputTokenDTO struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type ValidateTokenDTO struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
