package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionMode(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ActionMode
		output string
	}{
		{input: adaptivecard.ActionModePrimary, output: "primary"},
		{input: adaptivecard.ActionModeSecondary, output: "secondary"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.ActionModeString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.ActionModeString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsAActionMode())
	}

	_, err := adaptivecard.ActionModeString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ActionModeValues(), []adaptivecard.ActionMode{adaptivecard.ActionModePrimary, adaptivecard.ActionModeSecondary})
	assert.ElementsMatch(t, adaptivecard.ActionModeStrings(), []string{"primary", "secondary"})
}
