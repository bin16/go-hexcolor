package hexcolor

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	type data struct {
		Color Color
	}

	var colors = []string{
		"#DCE0D9",
		"#31081F",
		"#6B0F1A",
		"#595959",
		"#808F85",
		"#123",
		"#1234",
		"#123456",
		"#12345678",
		"#F0F0F0",
		"#F0F0F0F0",
		"#0F0F0F",
		"#0F0F0F0F",
		"#aaa",
		"#aaaa",
		"#bbb",
		"#000",
		"#0000",
		"#000000",
		"#00000000",
	}

	for _, s := range colors {
		var (
			c1 = New(s)
			c2 = Parse(s)
		)

		assert.Equal(t, c1.NRGBA, c2)

		var d1 = data{c1}

		var raw, err1 = json.Marshal(d1)
		assert.NoError(t, err1)

		assert.True(t, bytes.Contains(
			raw,
			[]byte(`"`+ToString(c1.NRGBA)+`"`),
		))

		var d2 = data{}

		var err2 = json.Unmarshal(raw, &d2)
		assert.NoError(t, err2)

		assert.Equal(t, d1, d2)
	}
}

func TestUnmarshal(t *testing.T) {
	var badColors = []string{
		`""`,
		`#`,
		`123`,
		`#1Fx@`,
		`"123`,
		`123"`,
		`//////`,
	}

	for _, s := range badColors {
		var c1 = Color{}
		assert.NoError(t, c1.UnmarshalJSON([]byte(s)))
	}
}

func TestMarshal(t *testing.T) {
	var colors = []string{
		"#DCE0D9",
		"#31081F",
		"#6B0F1A",
		"#595959",
		"#808F85",
		"#123",
		"#1234",
		"#123456",
		"#12345678",
		"#F0F0F0",
		"#F0F0F0F0",
		"#0F0F0F",
		"#0F0F0F0F",
		"#aaa",
		"#aaaa",
		"#bbb",
		"#000",
		"#0000",
		"#000000",
		"#00000000",
	}

	for _, s := range colors {
		var c = New(s)
		var b1, err = c.MarshalJSON()
		assert.NoError(t, err, s)

		assert.Equal(t, `"`+c.String()+`"`, string(b1))

		b1 = bytes.Trim(b1, `"`)
		var c1 = parseHexColor(s)
		var c2 = parseBytes(b1)

		assert.Equal(t, c1, c2, "%v (%s) != %v (%s)\n", c1, s, c2, b1)
	}
}
