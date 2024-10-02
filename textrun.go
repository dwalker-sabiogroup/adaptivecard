package adaptivecard

type TextRun struct {
	Type          string       `json:"Type"`
	Text          string       `json:"text"`
	Color         Color        `json:"color"`
	FontType      FontType     `json:"fontType"`
	Highlight     bool         `json:"highlight"`
	IsSubtle      bool         `json:"isSubtle"`
	Italic        bool         `json:"italic"`
	SelectAction  SelectAction `json:"selectAction"`
	Size          FontSize     `json:"size"`
	Strikethrough bool         `json:"strikethrough"`
	Underline     bool         `json:"underline"`
	Weight        FontWeight   `json:"weight"`
}

func NewTextRun(txt string, opts ...func(*TextRun)) *TextRun {
	tr := &TextRun{
		Type: "TextRun",
		Text: txt,
	}

	for _, o := range opts {
		o(tr)
	}

	return tr
}

func WithTextRunColor(c Color) func(*TextRun) {
	return func(tr *TextRun) {
		tr.Color = c
	}
}

func WithTextRunFontType(ft FontType) func(*TextRun) {
	return func(tr *TextRun) {
		tr.FontType = ft
	}
}

func WithTextRunHighlight() func(*TextRun) {
	return func(tr *TextRun) {
		tr.Highlight = true
	}
}

func WithTextRunSubtle() func(*TextRun) {
	return func(tr *TextRun) {
		tr.IsSubtle = true
	}
}

func WithTextRunItalic() func(*TextRun) {
	return func(tr *TextRun) {
		tr.Italic = true
	}
}

func WithTextRunSelectAction(sa SelectAction) func(*TextRun) {
	return func(tr *TextRun) {
		tr.SelectAction = sa
	}
}

func WithTextRunSize(fs FontSize) func(*TextRun) {
	return func(tr *TextRun) {
		tr.Size = fs
	}
}

func WithTextRunStrikethrough() func(*TextRun) {
	return func(tr *TextRun) {
		tr.Strikethrough = true
	}
}

func WithTextRunUnderline() func(*TextRun) {
	return func(tr *TextRun) {
		tr.Underline = true
	}
}

func WithTextRunWeight(fw FontWeight) func(*TextRun) {
	return func(tr *TextRun) {
		tr.Weight = fw
	}
}
