package hexcolor

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	type colorData struct {
		Color Color `json:"color"`
	}

	var sl = []string{
		"#229896",
		"#371862",
		"#291",
		"#8901",
		"#1987a706",
		"#279ff281",
		"123",
		"#120938112310-923",
	}

	for _, s := range sl {
		var c1 = New(s)
		var c2 = Parse(s)

		var r1, g1, b1, a1 = c1.RGBA()
		var r2, g2, b2, a2 = c2.RGBA()

		var n1 = parseHexColor(s)
		var n2 = c1.Color()
		assert.Equal(t, n1, n2)

		assert.Equal(t,
			[]uint32{r1, g1, b1, a1},
			[]uint32{r2, g2, b2, a2},
		)

		var j1 = colorData{c1}
		var bt, err = json.Marshal(j1)
		assert.NoError(t, err)

		var j2 = colorData{}
		assert.NoError(t, json.Unmarshal(bt, &j2))

		assert.Equal(t, j1, j2)
	}
}

func TestInvalidColors(t *testing.T) {
	var sl = []string{
		"",
		"1",
		"12",
		"12345",
		"1234567",
		"123456789",
		"(*)UY()09i011213",
	}

	for _, s := range sl {
		var c1 = New(s)

		assert.Equal(t, c1.Color(), Default, s)
	}
}
