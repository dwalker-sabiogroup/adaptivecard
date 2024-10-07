package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestTextRunDefault(t *testing.T) {
	tr := adaptivecard.NewTextRun("text")
	assert.Equal(t, "TextRun", tr.Type)
	assert.Equal(t, "text", tr.Text)
	assert.Equal(t, adaptivecard.ColorDefault, tr.Color)
	assert.Equal(t, adaptivecard.FontTypeDefault, tr.FontType)
	assert.False(t, tr.Highlight)
	assert.False(t, tr.IsSubtle)
	assert.False(t, tr.Italic)
	assert.Equal(t, adaptivecard.FontSizeDefault, tr.Size)
	assert.False(t, tr.Strikethrough)
	assert.False(t, tr.Underline)
	assert.Equal(t, adaptivecard.FontWeightDefault, tr.Weight)
}

func TestTextRunWithColor(t *testing.T) {
	assert.Equal(t, adaptivecard.ColorGood, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunColor(adaptivecard.ColorGood),
	).Color)
}

func TestTextRunWithFontType(t *testing.T) {
	assert.Equal(t, adaptivecard.FontTypeMonospace, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunFontType(adaptivecard.FontTypeMonospace),
	).FontType)
}

func TestTextRunWithHighlight(t *testing.T) {
	assert.True(t, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunHighlight(),
	).Highlight)
}

func TestTextRunWithSubtle(t *testing.T) {
	assert.True(t, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunSubtle(),
	).IsSubtle)
}

func TestTextRunWithItalic(t *testing.T) {
	assert.True(t, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunItalic(),
	).Italic)
}

func TestTextRunSize(t *testing.T) {
	assert.Equal(t, adaptivecard.FontSizeExtraLarge, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunSize(adaptivecard.FontSizeExtraLarge),
	).Size)
}

func TestTextRunWithStrikethrough(t *testing.T) {
	assert.True(t, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunStrikethrough(),
	).Strikethrough)
}

func TestTextRunWithUnderline(t *testing.T) {
	assert.True(t, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunUnderline(),
	).Underline)
}

func TestTextRunWithWeight(t *testing.T) {
	assert.Equal(t, adaptivecard.FontWeightBolder, adaptivecard.NewTextRun("text",
		adaptivecard.WithTextRunWeight(adaptivecard.FontWeightBolder),
	).Weight)
}
