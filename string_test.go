package rando

import (
	"fmt"
	"testing"
)

func TestStringsQtyLen(t *testing.T) {
	var c int
	for qty := 0; qty <= 260; qty += 26 {
		for length := 0; length <= 260; length += 26 {
			t.Run(fmt.Sprintf("test[%d]", c), func(t *testing.T) {
				got := StringsQtyLen(qty, length)
				if len(got) != qty {
					t.Errorf("len(StringsQtyLenSeed(%d, %d)) = %d, want %v", qty, length, len(got), qty)
				}

				for n := range got {
					if len(got[n]) != length {
						t.Errorf("len(got[n]) = %d, want %v", len(got[n]), length)
					}
				}
			})
			c++
		}
	}
}

func TestStringsQtyLenSeed(t *testing.T) {
	var c int
	for qty := 0; qty <= 260; qty += 26 {
		for length := 0; length <= 260; length += 26 {
			t.Run(fmt.Sprintf("test[%d]", c), func(t *testing.T) {
				got, gotSeed := StringsQtyLenSeed(qty, length)
				if gotSeed <= 0 {
					t.Errorf("gotSeed %d", gotSeed)
				}

				if len(got) != qty {
					t.Errorf("len(StringsQtyLenSeed(%d, %d)) = %d, want %v", qty, length, len(got), qty)
				}
				for n := range got {
					if len(got[n]) != length {
						t.Errorf("len(got[n]) = %d, want %v", len(got[n]), length)
					}
				}

				duplicate, newSeed := StringsQtyLenSeed(qty, length, gotSeed)
				if gotSeed != newSeed {
					t.Errorf("gotSeed %d != newSeed %d", gotSeed, newSeed)
				}
				if len(duplicate) != qty {
					t.Errorf("len(StringsQtyLenSeed(%d, %d, %d)) = %d, want %v", qty, length, gotSeed, len(duplicate), qty)
				}
				for n := range got {
					if got[n] != duplicate[n] {
						t.Errorf("got[n] =\n%s\nduplicate[n] =\n%s", got[n], duplicate[n])
					}
				}
			})
			c++
		}
	}
}
