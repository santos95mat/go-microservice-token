package usecase

import (
	"time"

	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
	"github.com/santos95mat/go-microservice-token/internal/interfaces"
	"github.com/santos95mat/go-microservice-token/internal/util"
)

type TokenUsecase struct {
	TokenRepository interfaces.TokenRepository
}

func NewTokenUsecase(tokenRepository interfaces.TokenRepository) *TokenUsecase {
	return &TokenUsecase{TokenRepository: tokenRepository}
}

func (u *TokenUsecase) ExecuteCreate(input dto.CreateTokenDTO) (*dto.OutputTokenDTO, error) {
	randonToken := util.RandonToken{
		LowCaseQuantity:     1,
		UpCaseQuantity:      1,
		NumbersQuantity:     1,
		SpecialCharQuantity: 1,
	}

	tokenStr := randonToken.Generate()
	validation := time.Now().Add(time.Duration(input.Validation) * time.Minute)

	token := entity.NewToken(input.UserID, tokenStr, validation)
	err := u.TokenRepository.Create(token)

	if err != nil {
		return nil, err
	}

	return &dto.OutputTokenDTO{
		Token:  token.Token,
		Expire: token.Expire,
	}, nil
}

func (u *TokenUsecase) ExecuteValidate(input dto.ValidateTokenDTO) error {
	err := u.TokenRepository.Validate(input)

	if err != nil {
		return err
	}

	return nil
}
