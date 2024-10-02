package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestMediaSourceDefault(t *testing.T) {
	ms := adaptivecard.NewMediaSource("http://example.com")

	assert.Equal(t, "http://example.com", ms.URL)
}

func TestMediaSourceMimeType(t *testing.T) {
	assert.Equal(t, "video/mp4", adaptivecard.NewMediaSource("http://example.com",
		adaptivecard.WithMediaSourceMimeType("video/mp4"),
	).MimeType)
}
