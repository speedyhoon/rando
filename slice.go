package rando

import "math/rand"

// BytesNil returns nil if the random length equals zero.
func BytesNil() (b []byte) {
	length := rand.Intn(100)
	for i := 0; i < length; i++ {
		b = append(b, Uint8())
	}
	return
}

// Bytes returns []byte{} if the random length equals zero.
func Bytes() (b []byte) {
	length := Uint8()
	b = make([]byte, length)
	for i := range b {
		b[i] = Uint8()
	}
	return
}
