package rando

import "math/rand"

// BytesNil returns nil if the random length equals zero.
func BytesNil() (b []byte) {
	length := rand.Intn(100)
	if length >= 1 {
		return BytesN(length)
	}
	return
}

// Bytes returns []byte{} if the random length equals zero.
func Bytes() (b []byte) {
	return BytesN(int(Uint8()))
}

// BytesN returns a fixed length slice filled with random bytes.
func BytesN(qty int) (b []byte) {
	b = make([]byte, qty)
	for i := 0; i < qty; i++ {
		b[i] = Uint8()
	}
	return
}
