package rando

import (
	"math/rand"
)

// Bool returns a random boolean.
func Bool() bool {
	return rand.Intn(2) == 1
}

// Bool10Percent returns a random boolean that's true roughly 10% of the time.
func Bool10Percent() bool {
	return rand.Intn(10) == 1
}
