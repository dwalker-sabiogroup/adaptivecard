package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestBlockElementHeight(t *testing.T) {
	tests := []struct {
		input  adaptivecard.BlockElementHeight
		output string
	}{
		{input: adaptivecard.BlockElementHeightAuto, output: "auto"},
		{input: adaptivecard.BlockElementHeightStretch, output: "stretch"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())
	}
}