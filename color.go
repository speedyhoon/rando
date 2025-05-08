package rando

import "github.com/speedyhoon/clr"

// C16 returns a random 16-bit color from 27 possible web-safe primary colors.
// Each red, green and blue channel uses either 0, 7 or F as their hexadecimal digits.
func C16() clr.C16 {
	const size = 3
	c := [size]clr.C16{0x0, 0x7, 0xF} // Options to randomly select color values from.
	return c[Uint8n(size)]<<12 | c[Uint8n(size)]<<8 | c[Uint8n(size)]<<4 | 0xf
}

// C32 returns a random 24-bit color with 100% opacity.
func C32() clr.C32 {
	const MaxUint24 = 1<<24 - 1
	return clr.C32(Uint32n(MaxUint24))<<8 | 0xff
}
