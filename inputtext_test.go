package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputTextDefault(t *testing.T) {
	it := adaptivecard.NewInputText("id")

	assert.Equal(t, "Input.Text", it.Type)
	assert.Equal(t, "id", it.ID)
	assert.False(t, it.IsMultiline)
	assert.Empty(t, it.MaxLength)
	assert.Empty(t, it.Placeholder)
	assert.Empty(t, it.Regex)
	assert.Equal(t, adaptivecard.TextInputStyleText, it.Style)
	assert.Empty(t, it.InlineAction)
	assert.Empty(t, it.Value)
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

func TestWithInputTextMultiline(t *testing.T) {
	assert.True(t, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextMultiline(),
	).IsMultiline)
}

func TestWithInputTextMaxLength(t *testing.T) {
	assert.Equal(t, 100, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextMaxLength(100),
	).MaxLength)
}

func TestWithInputTextPlaceholder(t *testing.T) {
	assert.Equal(t, "placeholder", adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextPlaceholder("placeholder"),
	).Placeholder)
}

func TestWithInputTextRegex(t *testing.T) {
	assert.Equal(t, "\\d+", adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextRegex("\\d+"),
	).Regex)
}

func TestWithInputTextStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.TextInputStylePassword, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextStyle(adaptivecard.TextInputStylePassword),
	).Style)
}

func TestWithInputTextValue(t *testing.T) {
	assert.Equal(t, "test", adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextValue("test"),
	).Value)
}

func TestWithInputTextErrorMessage(t *testing.T) {
	assert.Equal(t, "error", adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextErrorMessage("error"),
	).ErrorMessage)
}

func TestWithInputTextRequired(t *testing.T) {
	assert.True(t, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextRequired(),
	).IsRequired)
}

func TestWithInputTextLabel(t *testing.T) {
	assert.Equal(t, "label", adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextLabel("label"),
	).Label)
}

func TestWithInputTextLabelPosition(t *testing.T) {
	assert.Equal(t, adaptivecard.InputLabelPositionInline, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextLabelPosition(adaptivecard.InputLabelPositionInline),
	).LabelPosition)
}

func TestWithInputTextLabelWidth(t *testing.T) {
	assert.Equal(t, 10, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextLabelWidth(10),
	).LabelWidth)

	assert.Equal(t, "10px", adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextLabelWidth("10px"),
	).LabelWidth)
}

func TestWithInputTextInputStyle(t *testing.T) {
	assert.Equal(t, adaptivecard.InputStyleRevealOnHover, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextInputStyle(adaptivecard.InputStyleRevealOnHover),
	).InputStyle)
}

func TestWithInputTextHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextHeight(adaptivecard.BlockElementHeightStretch),
	).Height)
}

func TestWithInputTextSeparator(t *testing.T) {
	assert.True(t, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextSeparator(),
	).Separator)
}

func TestWithInputTextSpacing(t *testing.T) {
	assert.Equal(t, adaptivecard.SpacingMedium, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextSpacing(adaptivecard.SpacingMedium),
	).Spacing)
}

func TestWithInputTextNotVisible(t *testing.T) {
	assert.False(t, adaptivecard.NewInputText("id",
		adaptivecard.WithInputTextNotVisible(),
	).IsVisible)
}
