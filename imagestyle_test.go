package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ImageStyle
		output string
	}{
		{input: adaptivecard.ImageStyleDefault, output: "default"},
		{input: adaptivecard.ImageStylePerson, output: "person"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		is, err := adaptivecard.ImageStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, is)

		is, err = adaptivecard.ImageStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, is)
		assert.True(t, is.IsAImageStyle())
	}

	_, err := adaptivecard.ImageStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ImageStyleValues(), []adaptivecard.ImageStyle{adaptivecard.ImageStyleDefault, adaptivecard.ImageStylePerson})
	assert.ElementsMatch(t, adaptivecard.ImageStyleStrings(), []string{"person", "default"})
}
