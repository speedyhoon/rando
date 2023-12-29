package rando

import (
	"math/rand"
	"time"
)

// Bool returns a random boolean.
func Bool() bool {
	return rand.Intn(1) == 1
}

// DateTime returns a random year, month, day, hour, minute, second.
func DateTime() time.Time {
	return time.Unix(Int64(), 0).UTC()
}

// TimeNano returns a random year, month, day, hour, minute, second & nanosecond.
func TimeNano() time.Time {
	return time.Unix(0, Int64()).UTC()
}
