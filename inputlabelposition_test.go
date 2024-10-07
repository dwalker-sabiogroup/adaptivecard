package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputLabelPosition(t *testing.T) {
	tests := []struct {
		input  adaptivecard.InputLabelPosition
		output string
	}{
		{input: adaptivecard.InputLabelPositionInline, output: "inline"},
		{input: adaptivecard.InputLabelPositionAbove, output: "above"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ft, err := adaptivecard.InputLabelPositionString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)

		ft, err = adaptivecard.InputLabelPositionString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)
		assert.True(t, ft.IsAInputLabelPosition())
	}

	_, err := adaptivecard.InputLabelPositionString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.InputLabelPositionValues(), []adaptivecard.InputLabelPosition{adaptivecard.InputLabelPositionAbove, adaptivecard.InputLabelPositionInline})
	assert.ElementsMatch(t, adaptivecard.InputLabelPositionStrings(), []string{"inline", "above"})
}
