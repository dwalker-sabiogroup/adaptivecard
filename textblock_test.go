package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestTextBlockDefault(t *testing.T) {
	tb := adaptivecard.NewTextBlock("test")

	assert.Empty(t, tb.ID)
	assert.Equal(t, "test", tb.Text)
	assert.Equal(t, "TextBlock", tb.Type)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, tb.Height)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, tb.Height)
	assert.Equal(t, adaptivecard.ColorDefault, tb.Color)
	assert.Equal(t, adaptivecard.FontSizeDefault, tb.Size)
	assert.Equal(t, adaptivecard.FontTypeDefault, tb.FontType)
	assert.Equal(t, adaptivecard.HorizontalAlignmentLeft, tb.HorizontalAlignment)
	assert.Equal(t, adaptivecard.TextBlockStyleDefault, tb.Style)
	assert.False(t, tb.IsSubtle)
	assert.False(t, tb.Separator)
	assert.False(t, tb.Wrap)
	assert.True(t, tb.IsVisible)
	assert.Zero(t, tb.MaxLines)
}

func TestTextBlockOptions(t *testing.T) {
	tb := adaptivecard.NewTextBlock("test",
		adaptivecard.WithColor(adaptivecard.ColorAccent),
	)

	assert.Equal(t, adaptivecard.ColorAccent, tb.Color)
}
