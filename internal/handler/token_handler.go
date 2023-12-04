package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/santos95mat/go-microservice-token/internal/dto"
	"github.com/santos95mat/go-microservice-token/internal/interfaces"
	"github.com/santos95mat/go-microservice-token/pkg/mail"
)

type tokenHandler struct {
	tokenUsecase interfaces.TokenUsecase
}

func NewTokenHandler(tokenUsecase interfaces.TokenUsecase) *tokenHandler {
	return &tokenHandler{tokenUsecase: tokenUsecase}
}

func (h *tokenHandler) Create(c *fiber.Ctx) error {
	var input dto.CreateTokenDTO
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	_, err = uuid.Parse(input.UserID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	token, err := h.tokenUsecase.ExecuteCreate(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	go mail.SendMailHTML(token.Token, token.Expire, []string{input.Email})

	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *tokenHandler) Validate(c *fiber.Ctx) error {
	var input dto.ValidateTokenDTO
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	str, err := h.tokenUsecase.ExecuteValidate(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": str,
	})

}
