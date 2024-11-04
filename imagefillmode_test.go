package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageSetStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ImageSetStyle
		output string
	}{
		{input: adaptivecard.ImageSetStyleDefault, output: "default"},
		{input: adaptivecard.ImageSetStyleStacked, output: "stacked"},
		{input: adaptivecard.ImageSetStyleGrid, output: "grid"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ft, err := adaptivecard.ImageSetStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)

		ft, err = adaptivecard.ImageSetStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)
		assert.True(t, ft.IsAImageSetStyle())
	}

	_, err := adaptivecard.ImageSetStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ImageSetStyleValues(), []adaptivecard.ImageSetStyle{adaptivecard.ImageSetStyleDefault, adaptivecard.ImageSetStyleStacked, adaptivecard.ImageSetStyleGrid})
	assert.ElementsMatch(t, adaptivecard.ImageSetStyleStrings(), []string{"stacked", "default", "grid"})
}
