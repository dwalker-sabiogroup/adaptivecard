package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestActionSubmitDefault(t *testing.T) {
	action := adaptivecard.NewActionSubmit()

	assert.Equal(t, "Action.Submit", action.Type)
	assert.Len(t, action.Data, 0)
	assert.Equal(t, adaptivecard.AssociatedInputsAuto, action.AssociatedInputs)
	assert.Empty(t, action.Title)
	assert.Empty(t, action.IconUrl)
	assert.Empty(t, action.ID)
	assert.Equal(t, adaptivecard.ActionStyleDefault, action.Style)
	assert.Empty(t, action.Tooltip)
	assert.True(t, action.IsEnabled)
	assert.Equal(t, adaptivecard.ActionModePrimary, action.Mode)
}

func TestWithActionSubmitData(t *testing.T) {
	assert.Len(t, adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitData(map[string]interface{}{"key": "value"}),
	).Data, 1)

	assert.Equal(t, "value", adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitData(map[string]interface{}{"key": "value"}),
	).Data["key"])
}

func TestWithActionSubmitAssociatedInputs(t *testing.T) {
	assert.Equal(t, adaptivecard.AssociatedInputsNone, adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitAssociatedInputs(adaptivecard.AssociatedInputsNone),
	).AssociatedInputs)
}

func TestWithActionSubmitTitle(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitTitle("text"),
	).Title)
}

func TestWithActionSubmitIconURL(t *testing.T) {
	assert.Equal(t, "http://example.com", adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitIconURL("http://example.com"),
	).IconUrl)
}

func TestWithActionSubmitID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitID("id"),
	).ID)
}

func TestWithActionSubmitStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionStyleDestructive, adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitStyle(adaptivecard.ActionStyleDestructive),
	).Style)
}

func TestWithActionSubmitTooltip(t *testing.T) {
	assert.Equal(t, "tooltip", adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitTooltip("tooltip"),
	).Tooltip)
}

func TestWithActionSubmitDisabled(t *testing.T) {
	assert.False(t, adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitDisabled(),
	).IsEnabled)
}

func TestWithActionSubmitMode(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionModeSecondary, adaptivecard.NewActionSubmit(
		adaptivecard.WithActionSubmitMode(adaptivecard.ActionModeSecondary),
	).Mode)
}
