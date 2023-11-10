package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/santos95mat/go-microservice-token/internal/dto"
)

type TokenRepository interface {
	Create(token Token) error
	GetOne(search dto.TokenInputDTO) (Token, error)
}

type Token struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    string    `gorm:"not null"`
	Token     string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewToken(userID string, token string) Token {
	return Token{
		ID:     uuid.New(),
		UserID: userID,
		Token:  token,
	}
}
