package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssociatedInputs(t *testing.T) {
	tests := []struct {
		input  adaptivecard.AssociatedInputs
		output string
	}{
		{input: adaptivecard.AssociatedInputsAuto, output: "auto"},
		{input: adaptivecard.AssociatedInputsNone, output: "none"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.AssociatedInputsString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.AssociatedInputsString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsAAssociatedInputs())
	}

	_, err := adaptivecard.AssociatedInputsString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.AssociatedInputsValues(), []adaptivecard.AssociatedInputs{adaptivecard.AssociatedInputsAuto, adaptivecard.AssociatedInputsNone})
	assert.ElementsMatch(t, adaptivecard.AssociatedInputsStrings(), []string{"auto", "none"})
}
