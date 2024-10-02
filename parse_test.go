package hexcolor

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHexColor(t *testing.T) {
	var p = map[string]color.Color{
		"#000": color.NRGBA{0, 0, 0, 255},
		"#001": color.NRGBA{0, 0, 17, 255},
		"#010": color.NRGBA{0, 17, 0, 255},
		"#100": color.NRGBA{17, 0, 0, 255},
		"#112": color.NRGBA{17, 17, 34, 255},
		"#00f": color.NRGBA{0, 0, 255, 255},

		"#000000": color.NRGBA{0, 0, 0, 255},
		"#000001": color.NRGBA{0, 0, 1, 255},
		"#000010": color.NRGBA{0, 0, 16, 255},
		"#000011": color.NRGBA{0, 0, 17, 255},
		"#000100": color.NRGBA{0, 1, 0, 255},
		"#001000": color.NRGBA{0, 16, 0, 255},
		"#001100": color.NRGBA{0, 17, 0, 255},

		"":          Default,
		"1":         Default,
		"12":        Default,
		"123":       color.NRGBA{0x11, 0x22, 0x33, 0xff},
		"1234":      color.NRGBA{0x11, 0x22, 0x33, 0x44},
		"12345":     Default,
		"123456":    color.NRGBA{0x12, 0x34, 0x56, 0xff},
		"1234567":   Default,
		"12345678":  color.NRGBA{0x12, 0x34, 0x56, 0x78},
		"123456789": Default,

		"#":          Default,
		"#1":         Default,
		"#12":        Default,
		"#123":       color.NRGBA{0x11, 0x22, 0x33, 0xff},
		"#1234":      color.NRGBA{0x11, 0x22, 0x33, 0x44},
		"#12345":     Default,
		"#123456":    color.NRGBA{0x12, 0x34, 0x56, 0xff},
		"#1234567":   Default,
		"#12345678":  color.NRGBA{0x12, 0x34, 0x56, 0x78},
		"#123456789": Default,

		"#XX1234": color.NRGBA{0x00, 0x12, 0x34, 0xff},

		"0F0E0D":   color.NRGBA{15, 14, 13, 255},
		"#ffffff":  color.NRGBA{255, 255, 255, 255},
		"#CBE896":  color.NRGBA{203, 232, 150, 255},
		"#FFFFFC":  color.NRGBA{255, 255, 252, 255},
		"BEB7A4":   color.NRGBA{190, 183, 164, 255},
		"FF7F11":   color.NRGBA{255, 127, 17, 255},
		"#FF1B1C":  color.NRGBA{255, 27, 28, 255},
		"aBcDE012": color.NRGBA{0xab, 0xcd, 0xe0, 0x12},
	}

	for hc, c := range p {
		assert.Equal(t, c, Parse(hc), "<- %s", hc)
	}
}

func TestRToB(t *testing.T) {
	assert.Equal(t, byte(0), rToB('0'))
	assert.Equal(t, byte(1), rToB('1'))
	assert.Equal(t, byte(2), rToB('2'))
	assert.Equal(t, byte(3), rToB('3'))
	assert.Equal(t, byte(4), rToB('4'))
	assert.Equal(t, byte(5), rToB('5'))
	assert.Equal(t, byte(6), rToB('6'))
	assert.Equal(t, byte(7), rToB('7'))
	assert.Equal(t, byte(8), rToB('8'))
	assert.Equal(t, byte(9), rToB('9'))

	assert.Equal(t, byte(0xa), rToB('a'))
	assert.Equal(t, byte(0xb), rToB('b'))
	assert.Equal(t, byte(0xc), rToB('c'))
	assert.Equal(t, byte(0xd), rToB('d'))
	assert.Equal(t, byte(0xe), rToB('e'))
	assert.Equal(t, byte(0xf), rToB('f'))
	assert.Equal(t, byte(0xa), rToB('A'))
	assert.Equal(t, byte(0xb), rToB('B'))
	assert.Equal(t, byte(0xc), rToB('C'))
	assert.Equal(t, byte(0xd), rToB('D'))
	assert.Equal(t, byte(0xe), rToB('E'))
	assert.Equal(t, byte(0xf), rToB('F'))

	assert.Equal(t, byte(0x0), rToB('-'))
	assert.Equal(t, byte(0x0), rToB('X'))
	assert.Equal(t, byte(0x0), rToB('.'))
	assert.Equal(t, byte(0x0), rToB('&'))
	assert.Equal(t, byte(0x0), rToB(':'))
	assert.Equal(t, byte(0x0), rToB('\\'))
	assert.Equal(t, byte(0x0), rToB('P'))
	assert.Equal(t, byte(0x0), rToB('j'))
}

func BenchmarkRToB(b *testing.B) {
	var u = "0123456789abcdefABCDEF.x*&/="

	for i := 0; i < b.N; i++ {
		rToB(u[i%len(u)])
	}
}

func BenchmarkParse(b *testing.B) {
	var u = []string{
		"#012345",
		"#12345611",
		"#234567",
		"#34567833",
		"#456789",
		"#56789a55",
		"#6789aB",
		"#789AbC77",
		"#89aBcD",
		"#9AbCdE99",
		"#aBcDeF",
		"#bCdEf0bb",
		"#cDeF01",
		"#dEf012dd",
		"#eF0123",
		"#f01234ff",
		"012",
		"#123",
		"234",
		"#345",
		"567",
		"#789",
		"89A",
		"#9aB",
		"#ABC",
		"#bcd",
		"#dEF",
		"eFa",
		"FA0",
		"#A01",
	}

	for i := 0; i < b.N; i++ {
		Parse(u[i%len(u)])
	}
}
