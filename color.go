package hexcolor

import "image/color"

type Color string

func (clr Color) RGBA() (r, g, b, a uint32) {
	return Parse(string(clr)).RGBA()
}

func (clr Color) Color() color.Color {
	return Parse(string(clr))
}

func New(s string) Color {
	return Color(s)
}
