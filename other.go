package rando

import (
	"math/rand"
	"time"
)

// Bool returns a random boolean.
func Bool() bool {
	return rand.Intn(2) == 1
}

// Bool10Percent returns a random boolean that's true roughly 10% of the time.
func Bool10Percent() bool {
	return rand.Intn(10) == 1
}

// DateTime returns a random year, month, day, hour, minute, second.
func DateTime() time.Time {
	return time.Unix(Int64(), 0).UTC()
}

// TimeNano returns a random year, month, day, hour, minute, second & nanosecond.
func TimeNano() time.Time {
	return time.Unix(0, Int64()).UTC()
}
