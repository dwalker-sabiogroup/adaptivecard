package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColor(t *testing.T) {
	tests := []struct {
		input  adaptivecard.Color
		output string
	}{
		{input: adaptivecard.ColorDefault, output: "default"},
		{input: adaptivecard.ColorDark, output: "dark"},
		{input: adaptivecard.ColorLight, output: "light"},
		{input: adaptivecard.ColorAccent, output: "accent"},
		{input: adaptivecard.ColorGood, output: "good"},
		{input: adaptivecard.ColorWarning, output: "warning"},
		{input: adaptivecard.ColorAttention, output: "attention"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.ColorString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.ColorString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsAColor())
	}

	_, err := adaptivecard.ColorString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ColorValues(), []adaptivecard.Color{adaptivecard.ColorDefault, adaptivecard.ColorDark, adaptivecard.ColorLight, adaptivecard.ColorAccent, adaptivecard.ColorGood, adaptivecard.ColorWarning, adaptivecard.ColorAttention})
	assert.ElementsMatch(t, adaptivecard.ColorStrings(), []string{"dark", "default", "light", "accent", "good", "warning", "attention"})
}
