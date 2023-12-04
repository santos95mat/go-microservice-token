package repository

import (
	"time"

	"github.com/santos95mat/go-microservice-token/initializer"
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

// Structure for database conversation
type tokenRepository struct {
}

// Function to create anew tokenrepository
func NewTokenRepository() *tokenRepository {
	return &tokenRepository{}
}

// Method to save a token in DB
func (*tokenRepository) Create(token entity.Token) error {
	err := initializer.DB.Create(&token).Error

	return err
}

// Method to validate a token in DB
func (*tokenRepository) Validate(search dto.ValidateTokenDTO) (string, error) {
	var token entity.Token
	err := initializer.DB.Where("user_id = ? AND token = ?", search.UserID, search.Token).Find(&token).Error

	if err != nil {
		return "", err
	}

	if time.Now().After(token.Expire) {
		return "invalid", nil
	}

	return "valid", nil
}
