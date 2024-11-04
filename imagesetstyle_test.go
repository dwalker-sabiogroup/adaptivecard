package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageFillMode(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ImageFillMode
		output string
	}{
		{input: adaptivecard.ImageFillModeCover, output: "cover"},
		{input: adaptivecard.ImageFillModeRepeatHorizontally, output: "repeatHorizontally"},
		{input: adaptivecard.ImageFillModeRepeatVertically, output: "repeatVertically"},
		{input: adaptivecard.ImageFillModeRepeat, output: "repeat"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ft, err := adaptivecard.ImageFillModeString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)

		ft, err = adaptivecard.ImageFillModeString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)
		assert.True(t, ft.IsAImageFillMode())
	}

	_, err := adaptivecard.ImageFillModeString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ImageFillModeValues(), []adaptivecard.ImageFillMode{adaptivecard.ImageFillModeCover, adaptivecard.ImageFillModeRepeatHorizontally, adaptivecard.ImageFillModeRepeatVertically, adaptivecard.ImageFillModeRepeat})
	assert.ElementsMatch(t, adaptivecard.ImageFillModeStrings(), []string{"cover", "repeat", "repeatVertically", "repeatHorizontally"})
}
