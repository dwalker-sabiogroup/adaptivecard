package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerticalAlignmentType(t *testing.T) {
	tests := []struct {
		input  adaptivecard.VerticalAlignment
		output string
	}{
		{input: adaptivecard.VerticalAlignmentTop, output: "top"},
		{input: adaptivecard.VerticalAlignmentCenter, output: "center"},
		{input: adaptivecard.VerticalAlignmentBottom, output: "bottom"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ha, err := adaptivecard.VerticalAlignmentString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ha)

		ha, err = adaptivecard.VerticalAlignmentString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ha)
		assert.True(t, ha.IsAVerticalAlignment())
	}

	_, err := adaptivecard.VerticalAlignmentString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.VerticalAlignmentValues(), []adaptivecard.VerticalAlignment{adaptivecard.VerticalAlignmentTop, adaptivecard.VerticalAlignmentCenter, adaptivecard.VerticalAlignmentBottom})
	assert.ElementsMatch(t, adaptivecard.VerticalAlignmentStrings(), []string{"top", "bottom", "center"})
}
