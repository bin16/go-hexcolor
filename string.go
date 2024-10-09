package hexcolor

import "image/color"

func toBytes(clr color.NRGBA) []byte {
	if clr.A == 0xFF {
		return []byte{
			'#',
			bToR(clr.R / 16),
			bToR(clr.R % 16),
			bToR(clr.G / 16),
			bToR(clr.G % 16),
			bToR(clr.B / 16),
			bToR(clr.B % 16),
		}
	}

	return []byte{
		'#',
		bToR(clr.R / 16),
		bToR(clr.R % 16),
		bToR(clr.G / 16),
		bToR(clr.G % 16),
		bToR(clr.B / 16),
		bToR(clr.B % 16),
		bToR(clr.A / 16),
		bToR(clr.A % 16),
	}
}

func ToString(clr color.NRGBA) string {
	return string(toBytes(clr))
}
