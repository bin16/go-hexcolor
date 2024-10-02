package hexcolor

import (
	"image/color"
	"strings"
)

var Default = color.NRGBA{255, 0, 255, 255} // Magenta

func parseHexColor(s string) (clr color.NRGBA) {
	clr = Default

	s = strings.Trim(s, "# ")

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

func rToB(r uint8) uint8 {
	switch {
	case r >= 'a' && r <= 'f':
		return r - 'a' + 10
	case r >= 'A' && r <= 'F':
		return r - 'A' + 10
	case r >= '0' && r <= '9':
		return r - '0'
	}

	return 0x0
}

func Parse(s string) color.Color {
	return parseHexColor(s)
}
