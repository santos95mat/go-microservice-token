package dto

import "time"

// Data structure for creating a token
type CreateTokenDTO struct {
	UserID     string `json:"user_id"`
	Email      string `json:"email"`
	Validation int    `json:"validation"`
}

// Data structure for token response
type OutputTokenDTO struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// Data structure for validating a token
type ValidateTokenDTO struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
