package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestFontWeight(t *testing.T) {
	tests := []struct {
		input  adaptivecard.FontWeight
		output string
	}{
		{input: adaptivecard.FontWeightDefault, output: "default"},
		{input: adaptivecard.FontWeightLighter, output: "lighter"},
		{input: adaptivecard.FontWeightBolder, output: "bolder"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())
	}
}
