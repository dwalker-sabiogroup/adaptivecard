package adaptivecard

type Inline interface {
	string | TextRun
}

type RichTextBlock struct {
	Type                string              `json:"type"`
	Inlines             []interface{}       `json:"inlines"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment"`
	Height              BlockElementHeight  `json:"height"`
	Separator           bool                `json:"separator"`
	Spacing             Spacing             `json:"spacing"`
	ID                  string              `json:"id"`
	IsVisible           bool                `json:"isVisible"`
}

func NewRichTextBlock(opts ...func(*RichTextBlock)) *RichTextBlock {
	tr := &RichTextBlock{
		Type:      "RichTextBlock",
		Inlines:   []interface{}{},
		IsVisible: true,
	}

	for _, o := range opts {
		o(tr)
	}

	return tr
}

func WithRichTextBlockInline[T Inline](i T) func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.Inlines = append(rtb.Inlines, i)
	}
}

func WithRichTextBlockHorizontalAlignment(ha HorizontalAlignment) func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.HorizontalAlignment = ha
	}
}

func WithRichTextBlockHeight(h BlockElementHeight) func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.Height = h
	}
}

func WithRichTextBlockSeparator() func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.Separator = true
	}
}

func WithRichTextBlockSpacing(s Spacing) func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.Spacing = s
	}
}

func WithRichTextBlockID(i string) func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.ID = i
	}
}

func WithRichTextBlockNotVisible() func(*RichTextBlock) {
	return func(rtb *RichTextBlock) {
		rtb.IsVisible = false
	}
}
