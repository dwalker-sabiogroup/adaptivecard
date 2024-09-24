package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	tests := []struct {
		input  adaptivecard.Color
		output string
	}{
		{input: adaptivecard.ColorDefault, output: "default"},
		{input: adaptivecard.ColorDark, output: "dark"},
		{input: adaptivecard.ColorLight, output: "light"},
		{input: adaptivecard.ColorAccent, output: "accent"},
		{input: adaptivecard.ColorGood, output: "good"},
		{input: adaptivecard.ColorWarning, output: "warning"},
		{input: adaptivecard.ColorAttention, output: "attention"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())
	}
}
