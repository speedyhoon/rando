package rando

import "math/rand"

// Float32s returns a random length slice containing between 1 and 30 float32 numbers.
func Float32s() (f []float32) {
	l := rand.Intn(29) + 1
	f = make([]float32, l)
	for i := 0; i < l; i++ {
		f[i] = Float32()
	}
	return
}

// Float64s returns a random length slice containing between 1 and 30 float64 numbers.
func Float64s() (f []float64) {
	l := rand.Intn(29) + 1
	f = make([]float64, l)
	for i := 0; i < l; i++ {
		f[i] = Float64()
	}
	return
}
