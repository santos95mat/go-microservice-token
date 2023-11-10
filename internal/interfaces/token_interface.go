package interfaces

import (
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

type TokenRepository interface {
	Create(token entity.Token) error
	Validate(search dto.TokenInputDTO) (entity.Token, error)
}

type TokenUsecase interface {
	ExecuteCreate(input dto.TokenInputDTO) (dto.TokenOutputDTO, error)
	ExecuteValidate(search dto.TokenInputDTO) (dto.TokenOutputDTO, error)
}
