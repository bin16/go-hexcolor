package hexcolor

import (
	"image/color"
)

var (
	Default = color.NRGBA{}
)

const (
	quote = '"'
	hash  = '#'
)

func parseBytes(s []byte) (clr color.NRGBA) {
	clr = Default

	if len(s) < 3 {
		return
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) == 3 || len(s) == 6 {
		clr.A = 0xFF
	}

	if len(s) == 8 {
		clr.A = (rToB(s[6]) << 4) + rToB(s[7])
	}

	if len(s) == 6 || len(s) == 8 {
		clr.R = (rToB(s[0]) << 4) + rToB(s[1])
		clr.G = (rToB(s[2]) << 4) + rToB(s[3])
		clr.B = (rToB(s[4]) << 4) + rToB(s[5])

		return
	}

	if len(s) == 4 {
		clr.A = rToB(s[3]) * 17
	}

	if len(s) == 3 || len(s) == 4 {
		clr.R = rToB(s[0]) * 17
		clr.G = rToB(s[1]) * 17
		clr.B = rToB(s[2]) * 17

		return
	}

	return
}

func parseHexColor(s string) (clr color.NRGBA) {
	return parseBytes([]byte(s))
}

func Parse(s string) color.NRGBA {
	return parseHexColor(s)
}

func SafeParse(s string) color.NRGBA {
	if IsValid(s) {
		return parseHexColor(s)
	}

	return Default
}
