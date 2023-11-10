package app

import (
	"github.com/gofiber/fiber/v2"
)

func addTokenRoutes(token fiber.Router) {
	token.Post("/Token", func(c *fiber.Ctx) error {
		return c.SendString("Create a token")
	})

	token.Get("/Token", func(c *fiber.Ctx) error {
		return c.SendString("Get a token")
	})
}
