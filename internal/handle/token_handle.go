package handle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	var input dto.CreateTokenDTO
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	_, err = uuid.Parse(input.UserID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	token, err := h.TokenUsecase.ExecuteCreate(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *tokenHandle) Validate(c *fiber.Ctx) error {
	var input dto.ValidateTokenDTO
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	str, err := h.TokenUsecase.ExecuteValidate(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": str,
	})

}
