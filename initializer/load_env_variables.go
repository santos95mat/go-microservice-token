package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

// Function that reads the project's environment variables
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
