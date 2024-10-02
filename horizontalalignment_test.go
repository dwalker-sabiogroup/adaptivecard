package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHorizontalAlignmentType(t *testing.T) {
	tests := []struct {
		input  adaptivecard.HorizontalAlignment
		output string
	}{
		{input: adaptivecard.HorizontalAlignmentLeft, output: "left"},
		{input: adaptivecard.HorizontalAlignmentCenter, output: "center"},
		{input: adaptivecard.HorizontalAlignmentRight, output: "right"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))
	}
}
