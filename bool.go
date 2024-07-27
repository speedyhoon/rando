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
	return BoolsN(Uint8n(100))
}

// BoolsN returns a []bool with length `size`, populated with random values.
func BoolsN(size uint8) (b []bool) {
	if size >= 1 {
		b = make([]bool, size)
		for i := range b {
			b[i] = Bool()
		}
	}

	return
}
