package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputNumberDefault(t *testing.T) {
	in := adaptivecard.NewInputNumber("id")

	assert.Equal(t, "Input.Number", in.Type)
	assert.Equal(t, "id", in.ID)
	assert.Empty(t, in.Max)
	assert.Empty(t, in.Min)
	assert.Empty(t, in.Placeholder)
	assert.Empty(t, in.Value)
	assert.Empty(t, in.ErrorMessage)
	assert.False(t, in.IsRequired)
	assert.Empty(t, in.Label)
	assert.Equal(t, adaptivecard.InputLabelPositionAbove, in.LabelPosition)
	assert.Equal(t, adaptivecard.InputStyleDefault, in.InputStyle)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, in.Height)
	assert.False(t, in.Separator)
	assert.Equal(t, adaptivecard.SpacingDefault, in.Spacing)
	assert.True(t, in.IsVisible)
}

func TestWithInputNumberMax(t *testing.T) {
	assert.Equal(t, 1, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberMax(1),
	).Max)
}

func TestWithInputNumberMin(t *testing.T) {
	assert.Equal(t, 1, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberMin(1),
	).Min)
}

func TestWithInputNumberPlaceholder(t *testing.T) {
	assert.Equal(t, "placeholder", adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberPlaceholder("placeholder"),
	).Placeholder)
}

func TestWithInputNumberValue(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberValue("test"),
	).Value)
}

func TestWithInputNumberErrorMessage(t *testing.T) {
	assert.Equal(t, "error", adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberErrorMessage("error"),
	).ErrorMessage)
}

func TestWithInputNumberRequired(t *testing.T) {
	assert.True(t, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberRequired(),
	).IsRequired)
}

func TestWithInputNumberLabel(t *testing.T) {
	assert.Equal(t, "label", adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberLabel("label"),
	).Label)
}

func TestWithInputNumberLabelPosition(t *testing.T) {
	assert.Equal(t, adaptivecard.InputLabelPositionInline, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberLabelPosition(adaptivecard.InputLabelPositionInline),
	).LabelPosition)
}

func TestWithInputNumberLabelWidth(t *testing.T) {
	assert.Equal(t, 10, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberLabelWidth(10),
	).LabelWidth)

	assert.Equal(t, "10px", adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberLabelWidth("10px"),
	).LabelWidth)
}

func TestWithInputNumberInputStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.InputStyleRevealOnHover, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberInputStyle(adaptivecard.InputStyleRevealOnHover),
	).InputStyle)
}

func TestWithInputNumberHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithInputNumberSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberSeparator(),
	).Separator)
}

func TestWithInputNumberSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingMedium, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberSpacing(adaptivecard.SpacingMedium),
	).Spacing)
}

func TestWithInputNumberNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewInputNumber("id",
		adaptivecard.WithInputNumberNotVisible(),
	).IsVisible)
}
