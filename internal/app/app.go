package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Variable to instantiate a new webb application
var app = fiber.New()

// Function to run the web application
func Run() {
	app.Use(cors.New())

	addTokenRoutes(app)

	err := app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}
