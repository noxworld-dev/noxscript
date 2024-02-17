package ns

import "image/color"

type Color = color.Color

// RGB creates an RGB color.
func RGB(r, g, b byte) Color {
	return color.NRGBA{R: r, G: g, B: b, A: 255}
}

// RGBA creates an RGBA color.
func RGBA(r, g, b, a byte) Color {
	return color.NRGBA{R: r, G: g, B: b, A: a}
}
