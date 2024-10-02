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

func WithTextBlockMaxLines(i int) func(*TextBlock) {
	return func(t *TextBlock) {
		t.MaxLines = i
	}
}

func WithTextBlockFontSize(fs FontSize) func(*TextBlock) {
	return func(t *TextBlock) {
		t.Size = fs
	}
}

func WithTextBlockFontWeight(fw FontWeight) func(*TextBlock) {
	return func(t *TextBlock) {
		t.Weight = fw
	}
}

func WithTextBlockWrap() func(*TextBlock) {
	return func(t *TextBlock) {
		t.Wrap = true
	}
}

func WithTextBlockStyle(ts TextBlockStyle) func(*TextBlock) {
	return func(t *TextBlock) {
		t.Style = ts
	}
}

func WithTextBlockHeight(bh BlockElementHeight) func(*TextBlock) {
	return func(t *TextBlock) {
		t.Height = bh
	}
}

func WithTextBlockSeparator() func(*TextBlock) {
	return func(t *TextBlock) {
		t.Separator = true
	}
}

func WithTextBlockSpacing(s Spacing) func(*TextBlock) {
	return func(t *TextBlock) {
		t.Spacing = s
	}
}

func WithTextBlockID(s string) func(*TextBlock) {
	return func(t *TextBlock) {
		t.ID = s
	}
}

func WithTextBlockNotVisible() func(*TextBlock) {
	return func(t *TextBlock) {
		t.IsVisible = false
	}
}
