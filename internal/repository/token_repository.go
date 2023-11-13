package repository

import (
	"errors"
	"time"

	"github.com/santos95mat/go-microservice-token/initializer"
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

type tokenRepository struct {
}

func NewTokenRepository() *tokenRepository {
	return &tokenRepository{}
}

func (*tokenRepository) Create(token entity.Token) error {
	err := initializer.DB.Create(&token).Error

	return err
}

func (*tokenRepository) Validate(search dto.ValidateTokenDTO) error {
	var token entity.Token
	err := initializer.DB.Where("user_id = ? AND token = ?", search.UserID, search.Token).First(&token).Error

	if err != nil {
		return err
	}

	if time.Now().After(token.Expire) {
		err = errors.New("Error: Token Invalid")
	}

	return err
}
