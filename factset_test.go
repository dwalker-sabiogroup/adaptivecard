package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestFactSetDefault(t *testing.T) {
	fs := adaptivecard.NewFactSet([]adaptivecard.Fact{*adaptivecard.NewFact("title", "value")})

	assert.Equal(t, "FactSet", fs.Type)
	assert.Len(t, fs.Facts, 1)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, fs.Height)
	assert.False(t, fs.Separator)
	assert.Empty(t, fs.Spacing)
	assert.Empty(t, fs.ID)
	assert.True(t, fs.IsVisible)
}

func TestFactSetSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewFactSet([]adaptivecard.Fact{*adaptivecard.NewFact("title", "value")},
		adaptivecard.WithFactSetSeparator(),
	).Separator)
}

func TestFactSetHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewFactSet([]adaptivecard.Fact{*adaptivecard.NewFact("title", "value")},
		adaptivecard.WithFactSetHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestFactSetSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewFactSet([]adaptivecard.Fact{*adaptivecard.NewFact("title", "value")},
		adaptivecard.WithFactSetSpacing(adaptivecard.SpacingExtraLarge),
	).Spacing)
}

func TestFactSetID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewFactSet([]adaptivecard.Fact{*adaptivecard.NewFact("title", "value")},
		adaptivecard.WithFactSetID("id"),
	).ID)
}

func TestFactSetNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewFactSet([]adaptivecard.Fact{*adaptivecard.NewFact("title", "value")},
		adaptivecard.WithFactSetNotVisible(),
	).IsVisible)
}
