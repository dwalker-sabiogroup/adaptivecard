package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectAciton(t *testing.T) {
	tests := []struct {
		input  adaptivecard.SelectAction
		output string
	}{
		{input: adaptivecard.SelectActionExecute, output: "Action.Execute"},
		{input: adaptivecard.SelectActionOpenURL, output: "Action.OpenUrl"},
		{input: adaptivecard.SelectActionSubmit, output: "Action.Submit"},
		{input: adaptivecard.SelectActionToggleVisibility, output: "Action.ToggleVisibility"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))
	}
}
