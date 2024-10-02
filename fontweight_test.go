package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFontWeight(t *testing.T) {
	tests := []struct {
		input  adaptivecard.FontWeight
		output string
	}{
		{input: adaptivecard.FontWeightDefault, output: "default"},
		{input: adaptivecard.FontWeightLighter, output: "lighter"},
		{input: adaptivecard.FontWeightBolder, output: "bolder"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ft, err := adaptivecard.FontWeightString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)

		ft, err = adaptivecard.FontWeightString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)
		assert.True(t, ft.IsAFontWeight())
	}

	_, err := adaptivecard.FontWeightString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.FontWeightValues(), []adaptivecard.FontWeight{adaptivecard.FontWeightDefault, adaptivecard.FontWeightLighter, adaptivecard.FontWeightBolder})
	assert.ElementsMatch(t, adaptivecard.FontWeightStrings(), []string{"lighter", "default", "bolder"})
}
