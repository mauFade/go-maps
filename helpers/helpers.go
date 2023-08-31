package helpers

import (
	"math/rand"
)

func RandomNumber(number int) int {
	value := rand.Intn(number)

	return value
}
