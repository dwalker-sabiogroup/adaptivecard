package adaptivecard_test

import (
	"strings"
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTextInputStyle(t *testing.T) {
	tests := []struct {
		input  adaptivecard.TextInputStyle
		output string
	}{
		{input: adaptivecard.TextInputStyleText, output: "text"},
		{input: adaptivecard.TextInputStyleTel, output: "tel"},
		{input: adaptivecard.TextInputStyleUrl, output: "url"},
		{input: adaptivecard.TextInputStyleEmail, output: "email"},
		{input: adaptivecard.TextInputStylePassword, output: "password"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))

		c, err := adaptivecard.TextInputStyleString(tc.output)
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)

		c, err = adaptivecard.TextInputStyleString(strings.ToUpper(tc.output))
		require.NoError(t, err)
		assert.Equal(t, tc.input, c)
		assert.True(t, c.IsATextInputStyle())
	}

	_, err := adaptivecard.TextInputStyleString("fake")
	require.Error(t, err)

	assert.ElementsMatch(t, adaptivecard.TextInputStyleValues(), []adaptivecard.TextInputStyle{adaptivecard.TextInputStyleText, adaptivecard.TextInputStyleTel, adaptivecard.TextInputStyleUrl, adaptivecard.TextInputStyleEmail, adaptivecard.TextInputStylePassword})
	assert.ElementsMatch(t, adaptivecard.TextInputStyleStrings(), []string{"text", "tel", "url", "email", "password"})
}
