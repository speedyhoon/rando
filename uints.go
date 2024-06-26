package rando

import (
	"math/rand"
)

var DefaultMax = 29

func Uints() (u []uint) {
	l := rand.Intn(DefaultMax) + 1
	u = make([]uint, l)
	for i := 0; i < l; i++ {
		if intSize == 32 {
			u[i] = uint(Uint32())
		} else {
			u[i] = uint(Uint64())
		}
	}
	return
}

func UintsN(qty int) (u []uint) {
	u = make([]uint, qty)
	for i := 0; i < qty; i++ {
		u[i] = Uint()
	}
	return
}

func Uint8s() (u []uint8) {
	l := rand.Intn(DefaultMax) + 1
	u = make([]uint8, l)
	for i := 0; i < l; i++ {
		u[i] = Uint8()
	}
	return
}

func Uint8sN(qty int) (u []uint8) {
	u = make([]uint8, qty)
	for i := 0; i < qty; i++ {
		u[i] = Uint8()
	}
	return
}

func Uint16s() (u []uint16) {
	l := rand.Intn(DefaultMax) + 1
	u = make([]uint16, l)
	for i := 0; i < l; i++ {
		u[i] = Uint16()
	}
	return
}

func Uint16sN(qty int) (u []uint16) {
	u = make([]uint16, qty)
	for i := 0; i < qty; i++ {
		u[i] = Uint16()
	}
	return
}

func Uint32s() (u []uint32) {
	l := rand.Intn(DefaultMax) + 1
	u = make([]uint32, l)
	for i := 0; i < l; i++ {
		u[i] = Uint32()
	}
	return
}

func Uint32sN(qty int) (u []uint32) {
	u = make([]uint32, qty)
	for i := 0; i < qty; i++ {
		u[i] = Uint32()
	}
	return
}

func Uint64s() (u []uint64) {
	l := rand.Intn(DefaultMax) + 1
	u = make([]uint64, l)
	for i := 0; i < l; i++ {
		u[i] = Uint64()
	}
	return
}

func Uint64sN(qty int) (u []uint64) {
	u = make([]uint64, qty)
	for i := 0; i < qty; i++ {
		u[i] = Uint64()
	}
	return
}
