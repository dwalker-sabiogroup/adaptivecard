package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
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
	}
}
