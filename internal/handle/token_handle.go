package handle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/interfaces"
)

type tokenHandle struct {
	TokenUsecase interfaces.TokenUsecase
}

func NewTokenHandle(tokenUsecase interfaces.TokenUsecase) *tokenHandle {
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

	token, err := h.TokenUsecase.ExecuteValidate(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(token)

}
