package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHorizontalAlignmentType(t *testing.T) {
	tests := []struct {
		input  adaptivecard.HorizontalAlignment
		output string
	}{
		{input: adaptivecard.HorizontalAlignmentLeft, output: "left"},
		{input: adaptivecard.HorizontalAlignmentCenter, output: "center"},
		{input: adaptivecard.HorizontalAlignmentRight, output: "right"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ha, err := adaptivecard.HorizontalAlignmentString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ha)

		ha, err = adaptivecard.HorizontalAlignmentString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ha)
		assert.True(t, ha.IsAHorizontalAlignment())
	}

	_, err := adaptivecard.HorizontalAlignmentString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.HorizontalAlignmentValues(), []adaptivecard.HorizontalAlignment{adaptivecard.HorizontalAlignmentLeft, adaptivecard.HorizontalAlignmentCenter, adaptivecard.HorizontalAlignmentRight})
	assert.ElementsMatch(t, adaptivecard.HorizontalAlignmentStrings(), []string{"left", "right", "center"})
}
