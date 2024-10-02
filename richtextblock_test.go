package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestRichTextBlockDefault(t *testing.T) {
	rtb := adaptivecard.NewRichTextBlock()

	assert.Equal(t, "RichTextBlock", rtb.Type)
	assert.Len(t, rtb.Inlines, 0)
	assert.Equal(t, adaptivecard.HorizontalAlignmentLeft, rtb.HorizontalAlignment)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, rtb.Height)
	assert.False(t, rtb.Separator)
	assert.True(t, rtb.IsVisible)
}

func TestWithRichTextBlockInline(t *testing.T) {
	assert.Len(t, adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockInline("string"),
		adaptivecard.WithRichTextBlockInline(*adaptivecard.NewTextRun("textrun")),
	).Inlines, 2)
}

func TestWithRichTextBlockHorizontalAlignment(t *testing.T) {
	assert.Equal(t, adaptivecard.HorizontalAlignmentRight, adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockHorizontalAlignment(adaptivecard.HorizontalAlignmentRight),
	).HorizontalAlignment)
}

func TestWithRichTextBlockHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithRichTextBlockSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockSeparator(),
	).Separator)
}

func TestWithRichTextBlockSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockSpacing(adaptivecard.SpacingExtraLarge),
	).Spacing)
}

func TestWithRichTextBlockID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockID("id"),
	).ID)
}

func TestWithRichTextBlockNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewRichTextBlock(
		adaptivecard.WithRichTextBlockNotVisible(),
	).IsVisible)
}
