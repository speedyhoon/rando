package rando

import (
	"math/rand"
)

func Ints() (u []int) {
	l := rand.Intn(defaultMax) + 1
	u = make([]int, l)
	for i := 0; i < l; i++ {
		if intSize == 32 {
			u[i] = int(Int32())
		} else {
			u[i] = int(Int64())
		}
	}
	return
}

func Int8s() (u []int8) {
	l := rand.Intn(defaultMax) + 1
	u = make([]int8, l)
	for i := 0; i < l; i++ {
		u[i] = Int8()
	}
	return
}

func Int16s() (u []int16) {
	l := rand.Intn(defaultMax) + 1
	u = make([]int16, l)
	for i := 0; i < l; i++ {
		u[i] = Int16()
	}
	return
}

func Int32s() (u []int32) {
	l := rand.Intn(defaultMax) + 1
	u = make([]int32, l)
	for i := 0; i < l; i++ {
		u[i] = Int32()
	}
	return
}

func Int64s() (u []int64) {
	l := rand.Intn(defaultMax) + 1
	u = make([]int64, l)
	for i := 0; i < l; i++ {
		u[i] = Int64()
	}
	return
}
