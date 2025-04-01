package rando

import "github.com/speedyhoon/clr"

// C16 returns a random 16-bit color from 27 possible web-safe primary colors.
// Each red, green and blue channel uses either 0, 7 or F as their hexadecimal digits.
func C16() clr.C16 {
	c := [3]clr.C16{0x0, 0x7, 0xF} // Options to randomly select color values from.
	return c[Uint8n(3)]<<12 | c[Uint8n(3)]<<8 | c[Uint8n(3)]<<4 | 0xf
}
