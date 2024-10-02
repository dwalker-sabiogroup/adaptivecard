package adaptivecard_test

import (
	"strings"
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

		tbs, err := adaptivecard.TextBlockStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, tbs)

		tbs, err = adaptivecard.TextBlockStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, tbs)
		assert.True(t, tbs.IsATextBlockStyle())
	}

	_, err := adaptivecard.TextBlockStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.TextBlockStyleValues(), []adaptivecard.TextBlockStyle{adaptivecard.TextBlockStyleDefault, adaptivecard.TextBlockStyleHeading})
	assert.ElementsMatch(t, adaptivecard.TextBlockStyleStrings(), []string{"heading", "default"})
}
