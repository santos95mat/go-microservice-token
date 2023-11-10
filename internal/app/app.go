package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	app        = fiber.New()
	tokenGroup = app.Group("/token")
)

func Run() {
	app.Use(cors.New())

	addTokenRoutes(tokenGroup)

	err := app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}
