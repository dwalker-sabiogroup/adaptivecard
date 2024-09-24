package adaptivecard

type TextBlock struct {
	Type                string              `json:"type"`
	Text                string              `json:"text"`
	Color               Color               `json:"color"`
	FontType            FontType            `json:"fonttype"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment"`
	IsSubtle            bool                `json:"isSubtle"`
	MaxLines            int                 `json:"maxLines"`
	Size                FontSize            `json:"size"`
	Weight              FontWeight          `json:"weight"`
	Wrap                bool                `json:"wrap"`
	Style               TextBlockStyle      `json:"style"`
	Height              BlockElementHeight  `json:"height"`
	Separator           bool                `json:"separator"`
	Spacing             Spacing             `json:"spacing"`
	ID                  string              `json:"id"`
	IsVisible           bool                `json:"isVisible"`
}

func NewTextBlock(text string, opts ...func(*TextBlock)) *TextBlock {
	t := &TextBlock{
		Type:      "TextBlock",
		Text:      text,
		IsVisible: true,
	}

	for _, o := range opts {
		o(t)
	}

	return t
}

func WithTextBlockColor(c Color) func(*TextBlock) {
	return func(t *TextBlock) {
		t.Color = c
	}
}

func WithTextBlockFontType(f FontType) func(*TextBlock) {
	return func(t *TextBlock) {
		t.FontType = f
	}
}

func WithTextBlockHorizontalAlignment(h HorizontalAlignment) func(*TextBlock) {
	return func(t *TextBlock) {
		t.HorizontalAlignment = h
	}
}

func WithTextBlockSubtle() func(*TextBlock) {
	return func(t *TextBlock) {
		t.IsSubtle = true
	}
}
