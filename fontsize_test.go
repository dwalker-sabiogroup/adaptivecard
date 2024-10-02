package adaptivecard_test

import (
	"strings"
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

		fs, err := adaptivecard.FontSizeString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, fs)

		fs, err = adaptivecard.FontSizeString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, fs)
		assert.True(t, fs.IsAFontSize())

	}

	_, err := adaptivecard.FontSizeString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.FontSizeValues(), []adaptivecard.FontSize{adaptivecard.FontSizeDefault, adaptivecard.FontSizeExtraLarge, adaptivecard.FontSizeLarge, adaptivecard.FontSizeMedium, adaptivecard.FontSizeSmall})
	assert.ElementsMatch(t, adaptivecard.FontSizeStrings(), []string{"small", "default", "medium", "large", "extraLarge"})
}
