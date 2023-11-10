package entity

import (
	"time"

	"github.com/google/uuid"
)

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
