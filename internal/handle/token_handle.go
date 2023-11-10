package handle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/repository"
	"github.com/santos95mat/go-microservice-token/internal/usecase"
)

var (
	tokenRepository = repository.NewTokenRepository()
	tokenUsecase    = usecase.NewTokenUsecase(tokenRepository)
	TokenHandler    = newTokenHandle(tokenUsecase)
)

type tokenHandle struct {
	TokenUsecase *usecase.TokenUsecase
}

func newTokenHandle(tokenUsecase *usecase.TokenUsecase) *tokenHandle {
	return &tokenHandle{TokenUsecase: tokenUsecase}
}

func (h *tokenHandle) Create(c *fiber.Ctx) error {
	var input dto.TokenInputDTO
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	token, err := h.TokenUsecase.ExecuteCreate(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *tokenHandle) GetOne(c *fiber.Ctx) error {
	var input dto.TokenInputDTO

	q := c.Queries()
	input.UserID = q["userID"]
	input.Token = q["token"]

	token, err := h.TokenUsecase.ExecuteGetOne(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(token)

}
