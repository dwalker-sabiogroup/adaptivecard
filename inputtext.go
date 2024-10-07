package adaptivecard

type SelectAction interface {
	ActionExecute | ActionOpenURL | ActionSubmit | ActionToggleVisibility
}

type LabelWidth interface {
	int | string
}

type InputText struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	IsMultiline   bool               `json:"isMultiline"`
	MaxLength     int                `json:"maxLength"`
	Placeholder   string             `json:"placeholder"`
	Regex         string             `json:"regex"`
	Style         TextInputStyle     `json:"style"`
	InlineAction  interface{}        `json:"inlineAction"`
	Value         string             `json:"value"`
	ErrorMessage  string             `json:"errorMessage"`
	IsRequired    bool               `json:"isRequired"`
	Label         string             `json:"label"`
	LabelPosition InputLabelPosition `json:"labelPosition"`
	LabelWidth    interface{}        `json:"labelWidth"`
	InputStyle    InputStyle         `json:"inputStyle"`
	Height        BlockElementHeight `json:"height"`
	Separator     bool               `json:"separator"`
	Spacing       Spacing            `json:"spacing"`
	IsVisible     bool               `json:"isVisible"`
}

func NewInputText(id string, opts ...func(*InputText)) *InputText {
	it := &InputText{
		Type:      "Input.Text",
		ID:        id,
		IsVisible: true,
	}

	for _, o := range opts {
		o(it)
	}

	return it
}

func WithInputTextMultiline() func(*InputText) {
	return func(it *InputText) {
		it.IsMultiline = true
	}
}

func WithInputTextMaxLength(ml int) func(*InputText) {
	return func(it *InputText) {
		it.MaxLength = ml
	}
}

func WithInputTextPlaceholder(p string) func(*InputText) {
	return func(it *InputText) {
		it.Placeholder = p
	}
}

func WithInputTextRegex(r string) func(*InputText) {
	return func(it *InputText) {
		it.Regex = r
	}
}

func WithInputTextStyle(s TextInputStyle) func(*InputText) {
	return func(it *InputText) {
		it.Style = s
	}
}

func WithInputTextValue(v string) func(*InputText) {
	return func(it *InputText) {
		it.Value = v
	}
}

func WithInputTextErrorMessage(em string) func(*InputText) {
	return func(it *InputText) {
		it.ErrorMessage = em
	}
}

func WithInputTextRequired() func(*InputText) {
	return func(it *InputText) {
		it.IsRequired = true
	}
}

func WithInputTextLabel(l string) func(*InputText) {
	return func(it *InputText) {
		it.Label = l
	}
}

func WithInputTextLabelPosition(lp InputLabelPosition) func(*InputText) {
	return func(it *InputText) {
		it.LabelPosition = lp
	}
}

func WithInputTextLabelWidth[T LabelWidth](lw T) func(*InputText) {
	return func(it *InputText) {
		it.LabelWidth = lw
	}
}

func WithInputTextInputStyle(is InputStyle) func(*InputText) {
	return func(it *InputText) {
		it.InputStyle = is
	}
}

func WithInputTextHeight(beh BlockElementHeight) func(*InputText) {
	return func(it *InputText) {
		it.Height = beh
	}
}

func WithInputTextSeparator() func(*InputText) {
	return func(it *InputText) {
		it.Separator = true
	}
}

func WithInputTextSpacing(s Spacing) func(*InputText) {
	return func(it *InputText) {
		it.Spacing = s
	}
}

func WithInputTextNotVisible() func(*InputText) {
	return func(it *InputText) {
		it.IsVisible = false
	}
}
