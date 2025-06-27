package rando

import "math/rand"

func Complex64() complex64 {
	return complex(Float32(), Float32())
}

func Complex128() complex128 {
	return complex(Float64(), Float64())
}

// Complex64s returns a random length slice containing between 1 and 30 complex64 numbers.
func Complex64s() (f []complex64) {
	l := rand.Intn(DefaultMax) + 1
	if l == 0 {
		return nil
	}

	f = make([]complex64, l)
	for i := 0; i < l; i++ {
		f[i] = Complex64()
	}
	return
}

func Complex64sN(qty int) (u []complex64) {
	u = make([]complex64, qty)
	for i := 0; i < qty; i++ {
		u[i] = Complex64()
	}
	return
}

// Complex128s returns a random length slice containing between 1 and 30 complex128 numbers.
func Complex128s() (f []complex128) {
	l := rand.Intn(DefaultMax) + 1
	if l == 0 {
		return nil
	}

	f = make([]complex128, l)
	for i := 0; i < l; i++ {
		f[i] = Complex128()
	}
	return
}

func Complex128sN(qty int) (u []complex128) {
	u = make([]complex128, qty)
	for i := 0; i < qty; i++ {
		u[i] = Complex128()
	}
	return
}
