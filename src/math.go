package src

import (
	"math/rand"
	"time"
)

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
