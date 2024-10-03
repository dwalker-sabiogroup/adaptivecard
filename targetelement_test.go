package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestTargetElementDefault(t *testing.T) {
	te := adaptivecard.NewTargetElement("id")

	assert.Equal(t, "id", te.ElementID)
	assert.Empty(t, te.IsVisible)
}

func TestTargetElementVisible(t *testing.T) {
	assert.True(t, adaptivecard.NewTargetElement("test",
		adaptivecard.WithTargetElementVisible(),
	).IsVisible)
}

func TestTargetElementNotvisible(t *testing.T) {
	assert.False(t, adaptivecard.NewTargetElement("test",
		adaptivecard.WithTargetElementNotVisible(),
	).IsVisible)
}
