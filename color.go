package hexcolor

import (
	"image/color"
)

type Color struct {
	color.NRGBA
}

func New(s string) Color {
	return Color{
		NRGBA: parseHexColor(s),
	}
}

func (clr Color) MarshalJSON() (r []byte, err error) {
	r = []byte{quote}
	r = append(r, toBytes(clr.NRGBA)...)
	r = append(r, quote)
	return r, nil
}

// - ""
// - "#123"
// - "#1234"
// - "#123456"
// - "#123456FF"
// - invalid values
func (clr *Color) UnmarshalJSON(data []byte) error {
	if len(data) < 4 {
		return nil
	}

	if data[0] == quote {
		data = data[1:]
	}

	if data[0] == hash {
		data = data[1:]
	}

	if d := len(data) - 1; data[d] == quote {
		data = data[:d]
	}

	clr.NRGBA = parseBytes(data)
	return nil
}

func (clr Color) String() string {
	return ToString(clr.NRGBA)
}
