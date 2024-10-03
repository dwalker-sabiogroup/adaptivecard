package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestActionOpenURLDefault(t *testing.T) {
	action := adaptivecard.NewActionOpenURL("http://example.com")

	assert.Equal(t, "Action.OpenUrl", action.Type)
	assert.Equal(t, "http://example.com", action.URL)
	assert.Empty(t, action.Title)
	assert.Empty(t, action.IconUrl)
	assert.Empty(t, action.ID)
	assert.Equal(t, adaptivecard.ActionStyleDefault, action.Style)
	assert.Empty(t, action.Tooltip)
	assert.True(t, action.IsEnabled)
	assert.Equal(t, adaptivecard.ActionModePrimary, action.Mode)
}

func TestWithActionOpenURLTitle(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLTitle("text"),
	).Title)
}

func TestWithActionOpenURLIconURL(t *testing.T) {
	assert.Equal(t, "http://example.com", adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLIconURL("http://example.com"),
	).IconUrl)
}

func TestWithActionOpenURLID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLID("id"),
	).ID)
}

func TestWithActionOpenURLStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionStyleDestructive, adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLStyle(adaptivecard.ActionStyleDestructive),
	).Style)
}

func TestWithActionOpenURLTooltip(t *testing.T) {
	assert.Equal(t, "tooltip", adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLTooltip("tooltip"),
	).Tooltip)
}

func TestWithActionOpenURLDisabled(t *testing.T) {
	assert.False(t, adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLDisabled(),
	).IsEnabled)
}

func TestWithActionOpenURLMode(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionModeSecondary, adaptivecard.NewActionOpenURL("test",
		adaptivecard.WithActionOpenURLMode(adaptivecard.ActionModeSecondary),
	).Mode)
}
