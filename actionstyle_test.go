package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ActionStyle
		output string
	}{
		{input: adaptivecard.ActionStyleDefault, output: "default"},
		{input: adaptivecard.ActionStylePositive, output: "positive"},
		{input: adaptivecard.ActionStyleDestructive, output: "destructive"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.ActionStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.ActionStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsAActionStyle())
	}

	_, err := adaptivecard.ActionStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ActionStyleValues(), []adaptivecard.ActionStyle{adaptivecard.ActionStyleDefault, adaptivecard.ActionStylePositive, adaptivecard.ActionStyleDestructive})
	assert.ElementsMatch(t, adaptivecard.ActionStyleStrings(), []string{"default", "destructive", "positive"})
}
