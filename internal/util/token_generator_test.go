package util_test

import (
	"testing"

	"github.com/santos95mat/go-microservice-token/internal/util"
)

// Function to test a token generation
func testGenerate(t *testing.T) {
	tokenDto := util.RandonToken{
		LowCaseQuantity:     1,
		UpCaseQuantity:      1,
		NumbersQuantity:     1,
		SpecialCharQuantity: 1,
	}

	randonToken := tokenDto.Generate()

	if len(randonToken) != 4 {
		t.Error("O token recebido é diferente do esperado")
	}
}
