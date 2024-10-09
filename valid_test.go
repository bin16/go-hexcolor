package hexcolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {

	var goodColors = []string{
		"#123",
		"#1234",
		"#123456",
		"#12345678",
	}

	var badColors = []string{
		"#",
		"#1",
		"#12",
		"#12345",
		"#1234567",
		"#123456789",
		"#12345G",
		"#!@##$%",
	}

	for _, s := range goodColors {
		assert.True(t, IsValid(s), s)
	}

	for _, s := range badColors {
		assert.False(t, IsValid(s), s)
	}
}
