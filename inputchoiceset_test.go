package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputChoiceSetDefault(t *testing.T) {
	it := adaptivecard.NewInputChoiceSet("id")

	assert.Equal(t, "Input.ChoiceSet", it.Type)
	assert.Equal(t, "id", it.ID)
	assert.Len(t, it.Choices, 0)
	assert.Empty(t, it.ChoicesData)
	assert.False(t, it.IsMultiSelect)
	assert.Equal(t, adaptivecard.ChoiceInputStyleCompact, it.Style)
	assert.Empty(t, it.Value)
	assert.Empty(t, it.Placeholder)
	assert.False(t, it.Wrap)
	assert.Empty(t, it.ErrorMessage)
	assert.False(t, it.IsRequired)
	assert.Empty(t, it.Label)
	assert.Equal(t, adaptivecard.InputLabelPositionAbove, it.LabelPosition)
	assert.Equal(t, adaptivecard.InputStyleDefault, it.InputStyle)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, it.Height)
	assert.False(t, it.Separator)
	assert.Equal(t, adaptivecard.SpacingDefault, it.Spacing)
	assert.True(t, it.IsVisible)
}

func TestWithInputChoiceSetChoices(t *testing.T) {
	assert.Len(t, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetChoices([]adaptivecard.InputChoice{*adaptivecard.NewInputChoice("title", "value")}),
	).Choices, 1)
}

func TestWithInputChoiceSetDataQuery(t *testing.T) {
	assert.Equal(t, "dataset", adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetDataQuery(*adaptivecard.NewDataQuery("dataset")),
	).ChoicesData.Dataset)
}

func TestWithInputChoiceSetMultiSelect(t *testing.T) {
	assert.True(t, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetMultiSelect(),
	).IsMultiSelect)
}

func TestWithInputChoiceSetWrap(t *testing.T) {
	assert.True(t, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetWrap(),
	).Wrap)
}

func TestWithInputChoiceSetPlaceholder(t *testing.T) {
	assert.Equal(t, "placeholder", adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetPlaceholder("placeholder"),
	).Placeholder)
}

func TestWithInputChoiceSetStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.ChoiceInputStyleExpanded, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetStyle(adaptivecard.ChoiceInputStyleExpanded),
	).Style)
}

func TestWithInputChoiceSetValue(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetValue("test"),
	).Value)
}

func TestWithInputChoiceSetErrorMessage(t *testing.T) {
	assert.Equal(t, "error", adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetErrorMessage("error"),
	).ErrorMessage)
}

func TestWithInputChoiceSetRequired(t *testing.T) {
	assert.True(t, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetRequired(),
	).IsRequired)
}

func TestWithInputChoiceSetLabel(t *testing.T) {
	assert.Equal(t, "label", adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetLabel("label"),
	).Label)
}

func TestWithInputChoiceSetLabelPosition(t *testing.T) {
	assert.Equal(t, adaptivecard.InputLabelPositionInline, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetLabelPosition(adaptivecard.InputLabelPositionInline),
	).LabelPosition)
}

func TestWithInputChoiceSetLabelWidth(t *testing.T) {
	assert.Equal(t, 10, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetLabelWidth(10),
	).LabelWidth)

	assert.Equal(t, "10px", adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetLabelWidth("10px"),
	).LabelWidth)
}

func TestWithInputChoiceSetInputStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.InputStyleRevealOnHover, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetInputStyle(adaptivecard.InputStyleRevealOnHover),
	).InputStyle)
}

func TestWithInputChoiceSetHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithInputChoiceSetSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetSeparator(),
	).Separator)
}

func TestWithInputChoiceSetSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingMedium, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetSpacing(adaptivecard.SpacingMedium),
	).Spacing)
}

func TestWithInputChoiceSetNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewInputChoiceSet("id",
		adaptivecard.WithInputChoiceSetNotVisible(),
	).IsVisible)
}
