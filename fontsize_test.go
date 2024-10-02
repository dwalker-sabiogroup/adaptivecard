package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFontSize(t *testing.T) {
	tests := []struct {
		input  adaptivecard.FontSize
		output string
	}{
		{input: adaptivecard.FontSizeDefault, output: "default"},
		{input: adaptivecard.FontSizeSmall, output: "small"},
		{input: adaptivecard.FontSizeMedium, output: "medium"},
		{input: adaptivecard.FontSizeLarge, output: "large"},
		{input: adaptivecard.FontSizeExtraLarge, output: "extraLarge"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())
		assert.True(t, tc.input.IsAFontSize())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

	}
}
