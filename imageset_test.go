package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestImageSetDefault(t *testing.T) {
	is := adaptivecard.NewImageSet([]adaptivecard.Image{*adaptivecard.NewImage("http://example.com")})

	assert.Equal(t, "ImageSet", is.Type)
	assert.Len(t, is.Images, 1)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, is.Height)
	assert.Equal(t, adaptivecard.ImageSizeMedium, is.ImageSize)
	assert.False(t, is.Separator)
	assert.Empty(t, is.Spacing)
	assert.Empty(t, is.ID)
	assert.True(t, is.IsVisible)
}

func TestImageSetSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewImageSet([]adaptivecard.Image{*adaptivecard.NewImage("http://example.com")},
		adaptivecard.WithImageSetSeparator(),
	).Separator)
}

func TestImageSetHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewImageSet([]adaptivecard.Image{*adaptivecard.NewImage("http://example.com")},
		adaptivecard.WithImageSetHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestImageSetSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewImageSet([]adaptivecard.Image{*adaptivecard.NewImage("http://example.com")},
		adaptivecard.WithImageSetSpacing(adaptivecard.SpacingExtraLarge),
	).Spacing)
}

func TestImageSetID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewImageSet([]adaptivecard.Image{*adaptivecard.NewImage("http://example.com")},
		adaptivecard.WithImageSetID("id"),
	).ID)
}

func TestImageSetNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewImageSet([]adaptivecard.Image{*adaptivecard.NewImage("http://example.com")},
		adaptivecard.WithImageSetNotVisible(),
	).IsVisible)
}
