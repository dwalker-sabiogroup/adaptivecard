package adaptivecard

type InputToggle struct {
	Type          string             `json:"type"`
	Title         string             `json:"title"`
	ID            string             `json:"id"`
	Value         string             `json:"value"`
	ValueOff      string             `json:"valueOff"`
	ValueOn       string             `json:"valueOn"`
	Wrap          bool               `json:"wrap"`
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

func NewInputToggle(title, id string, opts ...func(*InputToggle)) *InputToggle {
	it := &InputToggle{
		Type:      "Input.Toggle",
		Title:     title,
		ID:        id,
		Value:     "false",
		ValueOff:  "false",
		ValueOn:   "true",
		IsVisible: true,
	}

	for _, o := range opts {
		o(it)
	}

	return it
}

func WithInputToggleValue(v string) func(*InputToggle) {
	return func(it *InputToggle) {
		it.Value = v
	}
}

func WithInputToggleValueOff(v string) func(*InputToggle) {
	return func(it *InputToggle) {
		it.ValueOff = v
	}
}

func WithInputToggleValueOn(v string) func(*InputToggle) {
	return func(it *InputToggle) {
		it.ValueOn = v
	}
}

func WithInputToggleWrap() func(*InputToggle) {
	return func(it *InputToggle) {
		it.Wrap = true
	}
}

func WithInputToggleErrorMessage(em string) func(*InputToggle) {
	return func(it *InputToggle) {
		it.ErrorMessage = em
	}
}

func WithInputToggleRequired() func(*InputToggle) {
	return func(it *InputToggle) {
		it.IsRequired = true
	}
}

func WithInputToggleLabel(l string) func(*InputToggle) {
	return func(it *InputToggle) {
		it.Label = l
	}
}

func WithInputToggleLabelPosition(lp InputLabelPosition) func(*InputToggle) {
	return func(it *InputToggle) {
		it.LabelPosition = lp
	}
}

func WithInputToggleLabelWidth[T LabelWidth](lw T) func(*InputToggle) {
	return func(it *InputToggle) {
		it.LabelWidth = lw
	}
}

func WithInputToggleInputStyle(is InputStyle) func(*InputToggle) {
	return func(it *InputToggle) {
		it.InputStyle = is
	}
}

func WithInputToggleHeight(beh BlockElementHeight) func(*InputToggle) {
	return func(it *InputToggle) {
		it.Height = beh
	}
}

func WithInputToggleSeparator() func(*InputToggle) {
	return func(it *InputToggle) {
		it.Separator = true
	}
}

func WithInputToggleSpacing(s Spacing) func(*InputToggle) {
	return func(it *InputToggle) {
		it.Spacing = s
	}
}

func WithInputToggleNotVisible() func(*InputToggle) {
	return func(it *InputToggle) {
		it.IsVisible = false
	}
}
