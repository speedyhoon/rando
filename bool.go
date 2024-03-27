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

// Bools returns a random length []bool populated with random values.
func Bools() (b []bool) {
	if u := Uint8n(100); u != 0 {
		b = make([]bool, u)
		for i := range b {
			b[i] = Bool()
		}
	}

	return
}
