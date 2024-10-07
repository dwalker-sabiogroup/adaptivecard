package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputDateDefault(t *testing.T) {
	id := adaptivecard.NewInputDate("id")

	assert.Equal(t, "Input.Date", id.Type)
	assert.Equal(t, "id", id.ID)
	assert.Empty(t, id.Max)
	assert.Empty(t, id.Min)
	assert.Empty(t, id.Placeholder)
	assert.Empty(t, id.Value)
	assert.Empty(t, id.ErrorMessage)
	assert.False(t, id.IsRequired)
	assert.Empty(t, id.Label)
	assert.Equal(t, adaptivecard.InputLabelPositionAbove, id.LabelPosition)
	assert.Equal(t, adaptivecard.InputStyleDefault, id.InputStyle)
	assert.Equal(t, adaptivecard.BlockElementHeightAuto, id.Height)
	assert.False(t, id.Separator)
	assert.Equal(t, adaptivecard.SpacingDefault, id.Spacing)
	assert.True(t, id.IsVisible)
}

func TestWithInputDateMax(t *testing.T) {
	assert.Equal(t, "2024-10-07", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateMax("2024-10-07"),
	).Max)
}

func TestWithInputDateMin(t *testing.T) {
	assert.Equal(t, "2024-10-07", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateMin("2024-10-07"),
	).Min)
}

func TestWithInputDatePlaceholder(t *testing.T) {
	assert.Equal(t, "placeholder", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDatePlaceholder("placeholder"),
	).Placeholder)
}

func TestWithInputDateValue(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateValue("test"),
	).Value)
}

func TestWithInputDateErrorMessage(t *testing.T) {
	assert.Equal(t, "error", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateErrorMessage("error"),
	).ErrorMessage)
}

func TestWithInputDateRequired(t *testing.T) {
	assert.True(t, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateRequired(),
	).IsRequired)
}

func TestWithInputDateLabel(t *testing.T) {
	assert.Equal(t, "label", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateLabel("label"),
	).Label)
}

func TestWithInputDateLabelPosition(t *testing.T) {
	assert.Equal(t, adaptivecard.InputLabelPositionInline, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateLabelPosition(adaptivecard.InputLabelPositionInline),
	).LabelPosition)
}

func TestWithInputDateLabelWidth(t *testing.T) {
	assert.Equal(t, 10, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateLabelWidth(10),
	).LabelWidth)

	assert.Equal(t, "10px", adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateLabelWidth("10px"),
	).LabelWidth)
}

func TestWithInputDateInputStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.InputStyleRevealOnHover, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateInputStyle(adaptivecard.InputStyleRevealOnHover),
	).InputStyle)
}

func TestWithInputDateHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithInputDateSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateSeparator(),
	).Separator)
}

func TestWithInputDateSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingMedium, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateSpacing(adaptivecard.SpacingMedium),
	).Spacing)
}

func TestWithInputDateNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewInputDate("id",
		adaptivecard.WithInputDateNotVisible(),
	).IsVisible)
}
