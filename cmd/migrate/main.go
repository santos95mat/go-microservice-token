package main

import (
	"github.com/santos95mat/go-microservice-token/initializer"
	"github.com/santos95mat/go-microservice-token/internal/entity"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
}

type token struct {
}

func main() {
	initializer.DB.AutoMigrate(&entity.Token{})
}
