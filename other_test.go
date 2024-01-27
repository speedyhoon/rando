package rando

import "testing"

func TestBool(t *testing.T) {
	tally := map[bool]uint{
		true:  0,
		false: 0,
	}
	for i := 0; i < 1000; i++ {
		tally[Bool()]++
	}

	if tally[true] == 0 || tally[false] == 0 {
		t.Errorf("Bool() didn't generate enough randomness true: %d, false: %d", tally[true], tally[false])
	}
}
