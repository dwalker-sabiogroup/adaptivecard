package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestImageDefault(t *testing.T) {
	image := adaptivecard.NewImage("test")

	assert.Equal(t, "Image", image.Type)
	assert.Equal(t, "test", image.URL)
	assert.Empty(t, image.AltText)
	assert.Empty(t, image.BackgroundColor)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, image.Height)
	assert.Equal(t, adaptivecard.HorizontalAlignmentLeft, image.HorizontalAlignment)
	assert.Empty(t, image.SelectAction)
	assert.Empty(t, image.Size)
	assert.Empty(t, image.Style)
	assert.Empty(t, image.Width)
	assert.False(t, image.Separator)
	assert.Empty(t, image.Spacing)
	assert.Empty(t, image.ID)
	assert.True(t, image.IsVisible)
}

func TestImageAltText(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewImage("test",
		adaptivecard.WithImageAltText("text"),
	).AltText)
}

func TestImageBackgroundColor(t *testing.T) {
	assert.Equal(t, "#00000", adaptivecard.NewImage("test",
		adaptivecard.WithImageBackgroupColor("#00000"),
	).BackgroundColor)
}

func TestImageHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewImage("test",
		adaptivecard.WithImageHeight(adaptivecard.BlockElementHeightStretch),
	).Height)

	assert.Equal(t, "50px", adaptivecard.NewImage("test",
		adaptivecard.WithImageHeight("50px"),
	).Height)
}

func TestImageHorizontalAlignment(t *testing.T) {
	assert.Equal(t, adaptivecard.HorizontalAlignmentCenter, adaptivecard.NewImage("test",
		adaptivecard.WithImageHorizontalAlignment(adaptivecard.HorizontalAlignmentCenter),
	).HorizontalAlignment)
}

func TestImageImageSize(t *testing.T) {
	assert.Equal(t, adaptivecard.ImageSizeLarge, adaptivecard.NewImage("test",
		adaptivecard.WithImageSize(adaptivecard.ImageSizeLarge),
	).Size)
}

func TestImageImageStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ImageStylePerson, adaptivecard.NewImage("test",
		adaptivecard.WithImageStyle(adaptivecard.ImageStylePerson),
	).Style)
}

func TestImageWidth(t *testing.T) {
	assert.Equal(t, "100px", adaptivecard.NewImage("test",
		adaptivecard.WithImageWidth("100px"),
	).Width)
}

func TestImageSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewImage("test",
		adaptivecard.WithImageSeparator(),
	).Separator)
}

func TestImageSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewImage("test",
		adaptivecard.WithImageSpacing(adaptivecard.SpacingExtraLarge),
	).Spacing)
}

func TestImageID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewImage("test",
		adaptivecard.WithImageID("id"),
	).ID)
}

func TestImageNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewImage("test",
		adaptivecard.WithImageNotVisible(),
	).IsVisible)
}
