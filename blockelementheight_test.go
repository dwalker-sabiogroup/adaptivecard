package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlockElementHeight(t *testing.T) {
	tests := []struct {
		input  adaptivecard.BlockElementHeight
		output string
	}{
		{input: adaptivecard.BlockElementHeightAuto, output: "auto"},
		{input: adaptivecard.BlockElementHeightStretch, output: "stretch"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		beh, err := adaptivecard.BlockElementHeightString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, beh)

		beh, err = adaptivecard.BlockElementHeightString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, beh)
		assert.True(t, beh.IsABlockElementHeight())
	}

	_, err := adaptivecard.BlockElementHeightString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.BlockElementHeightValues(), []adaptivecard.BlockElementHeight{adaptivecard.BlockElementHeightAuto, adaptivecard.BlockElementHeightStretch})
	assert.ElementsMatch(t, adaptivecard.BlockElementHeightStrings(), []string{"auto", "stretch"})
}
