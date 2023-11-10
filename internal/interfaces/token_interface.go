package interfaces

import (
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

type TokenRepository interface {
	Create(token entity.Token) error
	GetOne(search dto.TokenInputDTO) (entity.Token, error)
}
