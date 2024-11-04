package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestMetadataDefault(t *testing.T) {
	ms := adaptivecard.NewMetadata()

	assert.Empty(t, ms.WebURL)
}
func TestMetadataWebURL(t *testing.T) {
	assert.Equal(t, "http://example.com", adaptivecard.NewMetadata(
		adaptivecard.WithMetadataWebURL("http://example.com"),
	).WebURL)
}
