package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestBackgroundImageDefault(t *testing.T) {
	fs := adaptivecard.NewBackgroundImage("https://example.com")

	assert.Equal(t, "https://example.com", fs.URL)
	assert.Equal(t, adaptivecard.ImageFillModeCover, fs.FillMode)
	assert.Equal(t, adaptivecard.HorizontalAlignmentLeft, fs.HorizontalAlignment)
	assert.Equal(t, adaptivecard.VerticalAlignmentTop, fs.VerticalAlignment)
}

func TestBackgroundImageFillMode(t *testing.T) {
	assert.Equal(t, adaptivecard.ImageFillModeRepeat, adaptivecard.NewBackgroundImage("https://example.com",
		adaptivecard.WithBackgroundImageFillMode(adaptivecard.ImageFillModeRepeat),
	).FillMode)
}

func TestBackgroundImageVerticalAlignment(t *testing.T) {
	assert.Equal(t, adaptivecard.VerticalAlignmentBottom, adaptivecard.NewBackgroundImage("https://example.com",
		adaptivecard.WithBackgroundImageVerticalAlignment(adaptivecard.VerticalAlignmentBottom),
	).VerticalAlignment)
}

func TestBackgroundImageHorizontalAlignment(t *testing.T) {
	assert.Equal(t, adaptivecard.HorizontalAlignmentRight, adaptivecard.NewBackgroundImage("https://example.com",
		adaptivecard.WithBackgroundImageHorizontalAlignment(adaptivecard.HorizontalAlignmentRight),
	).HorizontalAlignment)
}
