package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestActionToggleVisibilityDefault(t *testing.T) {
	action := adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")})

	assert.Equal(t, "Action.ToggleVisibility", action.Type)
	assert.Len(t, action.TargetElements, 1)
	assert.Empty(t, action.Title)
	assert.Empty(t, action.IconUrl)
	assert.Empty(t, action.ID)
	assert.Equal(t, adaptivecard.ActionStyleDefault, action.Style)
	assert.Empty(t, action.Tooltip)
	assert.True(t, action.IsEnabled)
	assert.Equal(t, adaptivecard.ActionModePrimary, action.Mode)
}

func TestWithActionToggleVisibilityTitle(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityTitle("text"),
	).Title)
}

func TestWithActionToggleVisibilityIconURL(t *testing.T) {
	assert.Equal(t, "http://example.com", adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityIconURL("http://example.com"),
	).IconUrl)
}

func TestWithActionToggleVisibilityID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityID("id"),
	).ID)
}

func TestWithActionToggleVisibilityStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionStyleDestructive, adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityStyle(adaptivecard.ActionStyleDestructive),
	).Style)
}

func TestWithActionToggleVisibilityTooltip(t *testing.T) {
	assert.Equal(t, "tooltip", adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityTooltip("tooltip"),
	).Tooltip)
}

func TestWithActionToggleVisibilityDisabled(t *testing.T) {
	assert.False(t, adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityDisabled(),
	).IsEnabled)
}

func TestWithActionToggleVisibilityMode(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionModeSecondary, adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")},
		adaptivecard.WithActionToggleVisibilityMode(adaptivecard.ActionModeSecondary),
	).Mode)
}
