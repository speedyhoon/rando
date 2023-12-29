package rando

import "math/rand"

func Bytes() (b []byte) {
	length := rand.Intn(100)
	for i := 0; i < length; i++ {
		b = append(b, Uint8())
	}
	return
}
