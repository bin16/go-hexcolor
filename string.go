package hexcolor

import "image/color"

func ToString(clr color.NRGBA) string {
	if clr.A == 0xFF {
		return string([]byte{
			'#',
			bToR(clr.R / 16),
			bToR(clr.R % 16),
			bToR(clr.G / 16),
			bToR(clr.G % 16),
			bToR(clr.B / 16),
			bToR(clr.B % 16),
		})
	}

	return string([]byte{
		'#',
		bToR(clr.R / 16),
		bToR(clr.R % 16),
		bToR(clr.G / 16),
		bToR(clr.G % 16),
		bToR(clr.B / 16),
		bToR(clr.B % 16),
		bToR(clr.A / 16),
		bToR(clr.A % 16),
	})
}
