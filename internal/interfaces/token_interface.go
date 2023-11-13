package interfaces

import (
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

type TokenRepository interface {
	Create(token entity.Token) error
	Validate(input dto.ValidateTokenDTO) (string, error)
}

type TokenUsecase interface {
	ExecuteCreate(input dto.CreateTokenDTO) (*dto.OutputTokenDTO, error)
	ExecuteValidate(search dto.ValidateTokenDTO) (string, error)
}
