package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.InputStyle
		output string
	}{
		{input: adaptivecard.InputStyleDefault, output: "default"},
		{input: adaptivecard.InputStyleRevealOnHover, output: "revealOnHover"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.InputStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.InputStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsAInputStyle())
	}

	_, err := adaptivecard.InputStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.InputStyleValues(), []adaptivecard.InputStyle{adaptivecard.InputStyleDefault, adaptivecard.InputStyleRevealOnHover})
	assert.ElementsMatch(t, adaptivecard.InputStyleStrings(), []string{"default", "revealOnHover"})
}
