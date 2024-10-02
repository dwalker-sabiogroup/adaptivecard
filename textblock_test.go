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

func TestWithTextBlockColor(t *testing.T) {
	assert.Equal(t, adaptivecard.ColorAccent, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockColor(adaptivecard.ColorAccent)).Color,
	)
}

func TestWithTextBlockFontType(t *testing.T) {
	assert.Equal(t, adaptivecard.FontTypeMonospace, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockFontType(adaptivecard.FontTypeMonospace)).FontType,
	)
}

func TestWithTextBlockHorizontalAlignment(t *testing.T) {
	assert.Equal(t, adaptivecard.HorizontalAlignmentRight, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockHorizontalAlignment(adaptivecard.HorizontalAlignmentRight)).HorizontalAlignment,
	)
}

func TestWithTextBlockSubtle(t *testing.T) {
	assert.True(t, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockSubtle()).IsSubtle,
	)
}

func TestWithTextBlockMaxLines(t *testing.T) {
	assert.Equal(t, 100, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockMaxLines(100)).MaxLines,
	)
}

func TestWithTextBlockFontSize(t *testing.T) {
	assert.Equal(t, adaptivecard.FontSizeLarge, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockFontSize(adaptivecard.FontSizeLarge)).Size,
	)
}

func TestWithTextBlockFontWeight(t *testing.T) {
	assert.Equal(t, adaptivecard.FontWeightBolder, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockFontWeight(adaptivecard.FontWeightBolder)).Weight,
	)
}

func TestWithTextBlockWrap(t *testing.T) {
	assert.True(t, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockWrap()).Wrap,
	)
}

func TestWithTextBlockStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.TextBlockStyleHeading, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockStyle(adaptivecard.TextBlockStyleHeading)).Style,
	)
}

func TestWithTextBlockHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockHeight(adaptivecard.BlockElementHeightStretch)).Height,
	)
}

func TestWithTextBlockSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockSeparator()).Separator,
	)
}

func TestWithTextBlockSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockSpacing(adaptivecard.SpacingExtraLarge)).Spacing,
	)
}

func TestWithTextBlockID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockID("id")).ID,
	)
}

func TestWithTextBlockNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewTextBlock("test",
		adaptivecard.WithTextBlockNotVisible()).IsVisible,
	)
}
