package rando

import (
	"math/rand"
	"time"
)

func Duration(min, max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max-min))) + min
}
