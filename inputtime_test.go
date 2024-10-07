package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputTimeDefault(t *testing.T) {
	in := adaptivecard.NewInputTime("id")

	assert.Equal(t, "Input.Time", in.Type)
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

func TestWithInputTimeMax(t *testing.T) {
	assert.Equal(t, "13:23", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeMax("13:23"),
	).Max)
}

func TestWithInputTimeMin(t *testing.T) {
	assert.Equal(t, "13:23", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeMin("13:23"),
	).Min)
}

func TestWithInputTimePlaceholder(t *testing.T) {
	assert.Equal(t, "placeholder", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimePlaceholder("placeholder"),
	).Placeholder)
}

func TestWithInputTimeValue(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeValue("test"),
	).Value)
}

func TestWithInputTimeErrorMessage(t *testing.T) {
	assert.Equal(t, "error", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeErrorMessage("error"),
	).ErrorMessage)
}

func TestWithInputTimeRequired(t *testing.T) {
	assert.True(t, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeRequired(),
	).IsRequired)
}

func TestWithInputTimeLabel(t *testing.T) {
	assert.Equal(t, "label", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeLabel("label"),
	).Label)
}

func TestWithInputTimeLabelPosition(t *testing.T) {
	assert.Equal(t, adaptivecard.InputLabelPositionInline, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeLabelPosition(adaptivecard.InputLabelPositionInline),
	).LabelPosition)
}

func TestWithInputTimeLabelWidth(t *testing.T) {
	assert.Equal(t, 10, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeLabelWidth(10),
	).LabelWidth)

	assert.Equal(t, "10px", adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeLabelWidth("10px"),
	).LabelWidth)
}

func TestWithInputTimeInputStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.InputStyleRevealOnHover, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeInputStyle(adaptivecard.InputStyleRevealOnHover),
	).InputStyle)
}

func TestWithInputTimeHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithInputTimeSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeSeparator(),
	).Separator)
}

func TestWithInputTimeSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingMedium, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeSpacing(adaptivecard.SpacingMedium),
	).Spacing)
}

func TestWithInputTimeNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewInputTime("id",
		adaptivecard.WithInputTimeNotVisible(),
	).IsVisible)
}
