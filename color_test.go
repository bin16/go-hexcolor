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
	}

	for _, s := range sl {
		var c1 = New(s)
		var c2 = Parse(s)

		var r1, g1, b1, a1 = c1.RGBA()
		var r2, g2, b2, a2 = c2.RGBA()

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
