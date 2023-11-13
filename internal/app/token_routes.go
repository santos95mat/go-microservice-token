package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-microservice-token/internal/handle"
	"github.com/santos95mat/go-microservice-token/internal/repository"
	"github.com/santos95mat/go-microservice-token/internal/usecase"
)

var (
	tokenRepository = repository.NewTokenRepository()
	tokenUsecase    = usecase.NewTokenUsecase(tokenRepository)
	tokenHandle     = handle.NewTokenHandle(tokenUsecase)
)

func addTokenRoutes(token fiber.Router) {
	token.Post("/", tokenHandle.Create)

	token.Post("/validate", tokenHandle.Validate)
}
