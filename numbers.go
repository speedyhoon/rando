package rando

import (
	"math"
	"math/rand"
)

const intSize = 32 << (^uint(0) >> 63) // 32 or 64

// Byte returns a random byte.
func Byte() byte {
	return Uint8()
}

// Float32 returns a random float32.
func Float32() float32 {
	return rand.Float32() * float32(math.Pow(10, float64(rand.Intn(9)+1)))
}

// Float64 returns a random float64.
func Float64() float64 {
	return rand.Float64() * math.Pow(10, float64(rand.Intn(9)+1))
}

// Int returns a random integer whose values include negative & zero values.
func Int() int {
	if intSize == 32 {
		return int(Int32())
	}
	return int(Int64())
}

// Int8 returns a random int8 whose values include negative & zero values.
func Int8() int8 {
	return int8(Uint8() - math.MaxInt8 + 1)
}

// Int16 returns a random int16 whose values include negative & zero values.
func Int16() int16 {
	return int16(Uint16() - math.MaxInt16 + 1)
}

// Int32 returns a random int32 whose values include negative & zero values.
func Int32() int32 {
	return int32(Uint32() - math.MaxInt32 + 1)
}

// Int64 returns a random int64 whose values include negative & zero values.
func Int64() int64 {
	return int64(Uint64() - math.MaxInt64 + 1)
}

// Rune returns any random rune including null.
func Rune() rune {
	return Int32()
}

// Uint returns a random unsigned integer.
func Uint() uint {
	if intSize == 32 {
		return uint(Uint32())
	}
	return uint(Uint64())
}

// Uint8 returns a random uint8.
func Uint8() uint8 {
	return uint8(rand.Intn(math.MaxUint8))
}

// Uint16 returns a random uint16.
func Uint16() uint16 {
	return uint16(rand.Intn(math.MaxUint16))
}

// Uint32 returns a random uint32.
func Uint32() uint32 {
	return rand.Uint32()
}

// Uint64 returns a random uint64.
func Uint64() uint64 {
	return rand.Uint64()
}
