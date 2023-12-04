package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-microservice-token/internal/handler"
	"github.com/santos95mat/go-microservice-token/internal/repository"
	"github.com/santos95mat/go-microservice-token/internal/usecase"
)

var (
	tokenRepository = repository.NewTokenRepository()
	tokenUsecase    = usecase.NewTokenUsecase(tokenRepository)
	tokenHandler    = handler.NewTokenHandler(tokenUsecase)
)

func addTokenRoutes(token fiber.Router) {
	token.Post("/", tokenHandler.Create)

	token.Post("/validate", tokenHandler.Validate)
}
