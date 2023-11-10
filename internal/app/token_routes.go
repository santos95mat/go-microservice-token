package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-microservice-token/internal/handle"
)

func addTokenRoutes(token fiber.Router) {
	token.Post("/", handle.TokenHandler.Create)

	token.Get("/", handle.TokenHandler.GetOne)
}
