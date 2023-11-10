package repository

import (
	"github.com/santos95mat/go-microservice-token/initializer"
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

type tokenRepository struct {
}

func NewTokenRepository() *tokenRepository {
	return &tokenRepository{}
}

func (r *tokenRepository) Create(token entity.Token) error {
	err := initializer.DB.Create(&token).Error

	return err
}

func (r *tokenRepository) GetOne(search dto.TokenInputDTO) (entity.Token, error) {
	var token entity.Token
	err := initializer.DB.Where("user_id = ? AND token = ?", search.UserID, search.Token).First(&token).Error

	return token, err
}
