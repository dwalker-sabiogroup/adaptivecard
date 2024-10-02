package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageSize(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ImageSize
		output string
	}{
		{input: adaptivecard.ImageSizeAuto, output: "auto"},
		{input: adaptivecard.ImageSizeStretch, output: "stretch"},
		{input: adaptivecard.ImageSizeSmall, output: "small"},
		{input: adaptivecard.ImageSizeMedium, output: "medium"},
		{input: adaptivecard.ImageSizeLarge, output: "large"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		is, err := adaptivecard.ImageSizeString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, is)

		is, err = adaptivecard.ImageSizeString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, is)
		assert.True(t, is.IsAImageSize())
	}

	_, err := adaptivecard.ImageSizeString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ImageSizeValues(), []adaptivecard.ImageSize{adaptivecard.ImageSizeAuto, adaptivecard.ImageSizeStretch, adaptivecard.ImageSizeSmall, adaptivecard.ImageSizeMedium, adaptivecard.ImageSizeLarge})
	assert.ElementsMatch(t, adaptivecard.ImageSizeStrings(), []string{"auto", "small", "medium", "large", "stretch"})
}
