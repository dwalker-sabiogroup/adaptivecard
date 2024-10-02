package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestCaptionSourceDefault(t *testing.T) {
	cs := adaptivecard.NewCaptionSource("vtt", "http://example.com", "test-label")

	assert.Equal(t, "http://example.com", cs.URL)
	assert.Equal(t, "vtt", cs.MimeType)
	assert.Equal(t, "test-label", cs.Label)
}
