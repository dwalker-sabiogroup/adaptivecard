package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestActionExecuteDefault(t *testing.T) {
	action := adaptivecard.NewActionExecute()

	assert.Equal(t, "Action.Execute", action.Type)
	assert.Empty(t, action.Verb)
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

func TestWithActionExecuteVerb(t *testing.T) {
	assert.Equal(t, "post", adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteVerb("post"),
	).Verb)
}

func TestWithActionExecuteData(t *testing.T) {
	assert.Len(t, adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteData(map[string]interface{}{"key": "value"}),
	).Data, 1)

	assert.Equal(t, "value", adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteData(map[string]interface{}{"key": "value"}),
	).Data["key"])
}

func TestWithActionExecuteAssociatedInputs(t *testing.T) {
	assert.Equal(t, adaptivecard.AssociatedInputsNone, adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteAssociatedInputs(adaptivecard.AssociatedInputsNone),
	).AssociatedInputs)
}

func TestWithActionExecuteTitle(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteTitle("text"),
	).Title)
}

func TestWithActionExecuteIconURL(t *testing.T) {
	assert.Equal(t, "http://example.com", adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteIconURL("http://example.com"),
	).IconUrl)
}

func TestWithActionExecuteID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteID("id"),
	).ID)
}

func TestWithActionExecuteStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionStyleDestructive, adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteStyle(adaptivecard.ActionStyleDestructive),
	).Style)
}

func TestWithActionExecuteTooltip(t *testing.T) {
	assert.Equal(t, "tooltip", adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteTooltip("tooltip"),
	).Tooltip)
}

func TestWithActionExecuteDisabled(t *testing.T) {
	assert.False(t, adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteDisabled(),
	).IsEnabled)
}

func TestWithActionExecuteMode(t *testing.T) {
	assert.Equal(t, adaptivecard.ActionModeSecondary, adaptivecard.NewActionExecute(
		adaptivecard.WithActionExecuteMode(adaptivecard.ActionModeSecondary),
	).Mode)
}
