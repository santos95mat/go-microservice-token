package usecase

import (
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

type TokenUsecase struct {
	TokenRepository entity.TokenRepository
}

func NewTokenUsecase(tokenRepository entity.TokenRepository) TokenUsecase {
	return TokenUsecase{TokenRepository: tokenRepository}
}

func (u TokenUsecase) ExecuteCreate(input dto.TokenInputDTO) (dto.TokenOutputDTO, error) {
	token := entity.NewToken(input.UserID, input.Token)
	err := u.TokenRepository.Create(token)

	if err != nil {
		return dto.TokenOutputDTO{}, err
	}

	return dto.TokenOutputDTO{
		ID:     token.ID,
		UserID: token.UserID,
		Token:  token.Token,
	}, err
}

func (u TokenUsecase) ExecuteGetOne(search dto.TokenInputDTO) (dto.TokenOutputDTO, error) {
	token, err := u.TokenRepository.GetOne(search)

	if err != nil {
		return dto.TokenOutputDTO{}, err
	}

	return dto.TokenOutputDTO{
		ID:     token.ID,
		UserID: token.UserID,
		Token:  token.Token,
	}, err
}
