package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestMediaDefault(t *testing.T) {
	ms := adaptivecard.NewMediaSource("http://www.example.com")
	m := adaptivecard.NewMedia([]adaptivecard.MediaSource{*ms})

	assert.Equal(t, "Media", m.Type)
	assert.Len(t, m.Sources, 1)
	assert.Empty(t, m.Poster)
	assert.Empty(t, m.AltText)
	assert.Len(t, m.CaptionSources, 0)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, m.Height)
	assert.False(t, m.Separator)
	assert.Equal(t, adaptivecard.SpacingDefault, m.Spacing)
	assert.Empty(t, m.ID)
	assert.True(t, m.IsVisible)
}

func TestMediaPoster(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaPoster("text"),
	).Poster)
}

func TestMediaAltText(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaAltText("text"),
	).AltText)
}

func TestMediaCaptionSources(t *testing.T) {
	cs := adaptivecard.NewCaptionSource("test", "http://example.com", "test")
	assert.Len(t, adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaCaptionSources([]adaptivecard.CaptionSource{*cs}),
	).CaptionSources, 1)
}

func TestMediaHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestMediaSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaSeparator(),
	).Separator)
}

func TestMediaSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaSpacing(adaptivecard.SpacingExtraLarge),
	).Spacing)
}

func TestMediaID(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaID("text"),
	).ID)
}

func TestMediaNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewMedia([]adaptivecard.MediaSource{},
		adaptivecard.WithMediaNotVisible(),
	).IsVisible)
}
