package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestActionShowCardDefault(t *testing.T) {
	action := adaptivecard.NewActionShowCard()

	assert.Equal(t, "Action.ShowCard", action.Type)
	assert.Empty(t, action.Card)
	assert.Empty(t, action.Title)
	assert.Empty(t, action.IconUrl)
	assert.Empty(t, action.ID)
	assert.Equal(t, adaptivecard.ActionStyleDefault, action.Style)
	assert.Empty(t, action.Tooltip)
	assert.True(t, action.IsEnabled)
	assert.Equal(t, adaptivecard.ActionModePrimary, action.Mode)
}

func TestWithActionShowCardTitle(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardTitle("text"),
	).Title)
}

func TestWithActionShowCardIconURL(t *testing.T) {
	assert.Equal(t, "http://example.com", adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardIconURL("http://example.com"),
	).IconUrl)
}

func TestWithActionShowCardID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardID("id"),
	).ID)
}

func TestWithActionShowCardStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionStyleDestructive, adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardStyle(adaptivecard.ActionStyleDestructive),
	).Style)
}

func TestWithActionShowCardTooltip(t *testing.T) {
	assert.Equal(t, "tooltip", adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardTooltip("tooltip"),
	).Tooltip)
}

func TestWithActionShowCardDisabled(t *testing.T) {
	assert.False(t, adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardDisabled(),
	).IsEnabled)
}

func TestWithActionShowCardMode(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionModeSecondary, adaptivecard.NewActionShowCard(
		adaptivecard.WithActionShowCardMode(adaptivecard.ActionModeSecondary),
	).Mode)
}
