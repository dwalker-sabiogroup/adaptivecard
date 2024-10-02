package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTextBlockStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.TextBlockStyle
		output string
	}{
		{input: adaptivecard.TextBlockStyleDefault, output: "default"},
		{input: adaptivecard.TextBlockStyleHeading, output: "heading"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))
	}
}
