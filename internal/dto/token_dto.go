package dto

import "github.com/google/uuid"

type TokenInputDTO struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type TokenOutputDTO struct {
	ID     uuid.UUID `json:"id"`
	UserID string    `json:"user_id"`
	Token  string    `json:"token"`
}
