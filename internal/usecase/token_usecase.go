package usecase

import (
	"time"

	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/entity"
	"github.com/santos95mat/go-microservice-token/internal/interfaces"
	"github.com/santos95mat/go-microservice-token/internal/util"
)

// Structure for repository abstraction
type TokenUsecase struct {
	tokenRepository interfaces.TokenRepository
}

// Function to create a new tokenUseCase
func NewTokenUsecase(tokenRepository interfaces.TokenRepository) *TokenUsecase {
	return &TokenUsecase{tokenRepository: tokenRepository}
}

// Method to Execute Create repository function
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
	err := u.tokenRepository.Create(token)

	if err != nil {
		return nil, err
	}

	return &dto.OutputTokenDTO{
		Token:  token.Token,
		Expire: token.Expire,
	}, nil
}

// Method to Execute Validate repository function
func (u *TokenUsecase) ExecuteValidate(input dto.ValidateTokenDTO) (string, error) {
	str, err := u.tokenRepository.Validate(input)

	if err != nil {
		return str, err
	}

	return str, nil
}
