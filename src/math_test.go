package src

import (
	"testing"
)

func TestRandomNumber(t *testing.T) {
	n := RandomNumber(1, 10)

	if n < 1 && n > 10 {
		t.Fatalf("%d is not between 1 and 10.", n)
	}
}
