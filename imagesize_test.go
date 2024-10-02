package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImageSize(t *testing.T) {
	tests := []struct {
		input  adaptivecard.ImageSize
		output string
	}{
		{input: adaptivecard.ImageSizeAuto, output: "auto"},
		{input: adaptivecard.ImageSizeStretch, output: "stretch"},
		{input: adaptivecard.ImageSizeSmall, output: "small"},
		{input: adaptivecard.ImageSizeMedium, output: "medium"},
		{input: adaptivecard.ImageSizeLarge, output: "large"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.output, tc.input.String())

		out, err := tc.input.MarshalJSON()

		require.NoError(t, err)
		assert.Equal(t, `"`+tc.output+`"`, string(out))
	}
}
