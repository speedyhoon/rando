package rando

import "testing"

func TestBool(t *testing.T) {
	tally := map[bool]uint16{
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

func TestBool10Percent(t *testing.T) {
	tally := map[bool]uint16{
		true:  0,
		false: 0,
	}
	for i := 0; i < 1000; i++ {
		tally[Bool10Percent()]++
	}

	if tally[true] == 0 || tally[false] == 0 {
		t.Errorf("Bool10Percent() didn't generate enough randomness true: %d, false: %d", tally[true], tally[false])
	}

	percentage := float32(tally[false]) / float32(tally[true])
	if percentage < 7 || percentage > 13 {
		t.Errorf("Bool10Percent() == true %f%%, not ~10%%.", percentage)
	}
}
