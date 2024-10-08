package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestActionSetDefault(t *testing.T) {
	as := adaptivecard.NewActionSet()

	assert.Equal(t, "ActionSet", as.Type)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, as.Height)
	assert.False(t, as.Separator)
	assert.Empty(t, as.Spacing)
	assert.Empty(t, as.ID)
	assert.True(t, as.IsVisible)
}

func TestActionSetActions(t *testing.T) {
	assert.Len(t, adaptivecard.NewActionSet(
		adaptivecard.WithActionSetAction(*adaptivecard.NewActionExecute()),
		adaptivecard.WithActionSetAction(*adaptivecard.NewActionOpenURL("http://example.com")),
		adaptivecard.WithActionSetAction(*adaptivecard.NewActionSubmit()),
		adaptivecard.WithActionSetAction(*adaptivecard.NewActionToggleVisibility([]adaptivecard.TargetElement{*adaptivecard.NewTargetElement("id")})),
		adaptivecard.WithActionSetAction(*adaptivecard.NewActionShowCard()),
	).Actions, 5)
}

func TestActionSetSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewActionSet(
		adaptivecard.WithActionSetSeparator(),
	).Separator)
}

func TestActionSetHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewActionSet(
		adaptivecard.WithActionSetHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestActionSetSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingExtraLarge, adaptivecard.NewActionSet(
		adaptivecard.WithActionSetSpacing(adaptivecard.SpacingExtraLarge),
	).Spacing)
}

func TestActionSetID(t *testing.T) {
	assert.Equal(t, "id", adaptivecard.NewActionSet(
		adaptivecard.WithActionSetID("id"),
	).ID)
}

func TestActionSetNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewActionSet(
		adaptivecard.WithActionSetNotVisible(),
	).IsVisible)
}
