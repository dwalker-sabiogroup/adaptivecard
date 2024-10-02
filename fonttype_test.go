package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFontType(t *testing.T) {
	tests := []struct {
		input  adaptivecard.FontType
		output string
	}{
		{input: adaptivecard.FontTypeDefault, output: "default"},
		{input: adaptivecard.FontTypeMonospace, output: "monospace"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		ft, err := adaptivecard.FontTypeString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)

		ft, err = adaptivecard.FontTypeString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, ft)
		assert.True(t, ft.IsAFontType())
	}

	_, err := adaptivecard.FontTypeString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.FontTypeValues(), []adaptivecard.FontType{adaptivecard.FontTypeDefault, adaptivecard.FontTypeMonospace})
	assert.ElementsMatch(t, adaptivecard.FontTypeStrings(), []string{"monospace", "default"})
}
