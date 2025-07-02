package rando

import (
	"math/rand"
	"time"
)

func Duration() time.Duration {
	return time.Duration(Int64())
}

func DurationBetween(min, max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max-min))) + min
}

// Time returns a random year, month, day, hour, minute, second.
func Time() time.Time {
	return time.Unix(Int64(), 0).UTC()
}

// TimeNano returns a random year, month, day, hour, minute, second & nanosecond.
func TimeNano() time.Time {
	return time.Unix(0, Int64()).UTC()
}

func Times() (t []time.Time) {
	l := rand.Intn(DefaultMax) + 1
	if l == 0 {
		return nil
	}

	t = make([]time.Time, l)
	for i := 0; i < l; i++ {
		t[i] = Time()
	}
	return
}
