package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChoiceInputStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ChoiceInputStyle
		output string
	}{
		{input: adaptivecard.ChoiceInputStyleCompact, output: "compact"},
		{input: adaptivecard.ChoiceInputStyleExpanded, output: "expanded"},
		{input: adaptivecard.ChoiceInputStyleFiltered, output: "filtered"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.ChoiceInputStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.ChoiceInputStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsAChoiceInputStyle())
	}

	_, err := adaptivecard.ChoiceInputStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.ChoiceInputStyleValues(), []adaptivecard.ChoiceInputStyle{adaptivecard.ChoiceInputStyleCompact, adaptivecard.ChoiceInputStyleExpanded, adaptivecard.ChoiceInputStyleFiltered})
	assert.ElementsMatch(t, adaptivecard.ChoiceInputStyleStrings(), []string{"compact", "expanded", "filtered"})
}
