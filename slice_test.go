package rando

import (
	"reflect"
	"testing"
)

func TestBytes(t *testing.T) {
	var hasZero bool
	for i := 0; i < 1000; i++ {
		b := Bytes()
		if len(b) == 0 {
			hasZero = true
			if b == nil {
				t.Errorf("Bytes() = %+v, want <nil>", b)
			}
		}
	}

	if !hasZero {
		t.Error("Bytes() doesn't  create enough empty slices ([]byte{})")
	}
}

func TestBytesN(t *testing.T) {
	s := BytesN(0)
	empty := []byte{}
	if len(s) != 0 || s == nil || !reflect.DeepEqual(s, empty) {
		t.Errorf("BytesN(0) = %+v, want <nil>", s)
	}

	for i := 0; i < 256; i++ {
		b := BytesN(i)
		l := len(b)
		if l != i {
			t.Errorf("len(BytesNil()) = %d, want %d", l, i)
		}
	}
}

func TestBytesNil(t *testing.T) {
	var hasZero bool
	for i := 0; i < 1000; i++ {
		b := BytesNil()
		if len(b) == 0 {
			hasZero = true
			if b != nil {
				t.Errorf("BytesNil() = %+v, want <nil>", b)
			}
		}
	}

	if !hasZero {
		t.Error("BytesNil() doesn't  create enough <nil> slices")
	}
}
