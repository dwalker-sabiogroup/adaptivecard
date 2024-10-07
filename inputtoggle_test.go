package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputToggleDefault(t *testing.T) {
	it := adaptivecard.NewInputToggle("title", "id")

	assert.Equal(t, "Input.Toggle", it.Type)
	assert.Equal(t, "title", it.Title)
	assert.Equal(t, "id", it.ID)
	assert.Equal(t, "false", it.Value)
	assert.Equal(t, "false", it.ValueOff)
	assert.Equal(t, "true", it.ValueOn)
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

func TestWithInputToggleValue(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleValue("test"),
	).Value)
}

func TestWithInputToggleValueOn(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleValueOn("test"),
	).ValueOn)
}

func TestWithInputToggleValueOff(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleValueOff("test"),
	).ValueOff)
}

func TestWithInputToggleWrap(t *testing.T) {
	assert.True(t, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleWrap(),
	).Wrap)
}

func TestWithInputToggleErrorMessage(t *testing.T) {
	assert.Equal(t, "error", adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleErrorMessage("error"),
	).ErrorMessage)
}

func TestWithInputToggleRequired(t *testing.T) {
	assert.True(t, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleRequired(),
	).IsRequired)
}

func TestWithInputToggleLabel(t *testing.T) {
	assert.Equal(t, "label", adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleLabel("label"),
	).Label)
}

func TestWithInputToggleLabelPosition(t *testing.T) {
	assert.Equal(t, adaptivecard.InputLabelPositionInline, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleLabelPosition(adaptivecard.InputLabelPositionInline),
	).LabelPosition)
}

func TestWithInputToggleLabelWidth(t *testing.T) {
	assert.Equal(t, 10, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleLabelWidth(10),
	).LabelWidth)

	assert.Equal(t, "10px", adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleLabelWidth("10px"),
	).LabelWidth)
}

func TestWithInputToggleInputStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.InputStyleRevealOnHover, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleInputStyle(adaptivecard.InputStyleRevealOnHover),
	).InputStyle)
}

func TestWithInputToggleHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithInputToggleSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleSeparator(),
	).Separator)
}

func TestWithInputToggleSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingMedium, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleSpacing(adaptivecard.SpacingMedium),
	).Spacing)
}

func TestWithInputToggleNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewInputToggle("title", "id",
		adaptivecard.WithInputToggleNotVisible(),
	).IsVisible)
}
