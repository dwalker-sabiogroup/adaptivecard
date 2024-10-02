package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpacing(t *testing.T) {
	tests := []struct {
		input  adaptivecard.Spacing
		output string
	}{
		{input: adaptivecard.SpacingDefault, output: "default"},
		{input: adaptivecard.SpacingNone, output: "none"},
		{input: adaptivecard.SpacingSmall, output: "small"},
		{input: adaptivecard.SpacingMedium, output: "medium"},
		{input: adaptivecard.SpacingLarge, output: "large"},
		{input: adaptivecard.SpacingExtraLarge, output: "extraLarge"},
		{input: adaptivecard.SpacingPadding, output: "padding"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		s, err := adaptivecard.SpacingString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, s)

		s, err = adaptivecard.SpacingString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, s)
		assert.True(t, s.IsASpacing())
	}

	_, err := adaptivecard.SpacingString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.SpacingValues(), []adaptivecard.Spacing{adaptivecard.SpacingDefault, adaptivecard.SpacingNone, adaptivecard.SpacingSmall, adaptivecard.SpacingMedium, adaptivecard.SpacingLarge, adaptivecard.SpacingExtraLarge, adaptivecard.SpacingPadding})
	assert.ElementsMatch(t, adaptivecard.SpacingStrings(), []string{"none", "default", "small", "medium", "large", "extraLarge", "padding"})
}
