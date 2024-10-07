package adaptivecard

type InputDate struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Max           string             `json:"max"`
	Min           string             `json:"min"`
	Placeholder   string             `json:"placeholder"`
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

func NewInputDate(id string, opts ...func(*InputDate)) *InputDate {
	it := &InputDate{
		Type:      "Input.Date",
		ID:        id,
		IsVisible: true,
	}

	for _, o := range opts {
		o(it)
	}

	return it
}

func WithInputDateMax(m string) func(*InputDate) {
	return func(it *InputDate) {
		it.Max = m
	}
}

func WithInputDateMin(m string) func(*InputDate) {
	return func(it *InputDate) {
		it.Min = m
	}
}

func WithInputDatePlaceholder(p string) func(*InputDate) {
	return func(it *InputDate) {
		it.Placeholder = p
	}
}

func WithInputDateValue(v string) func(*InputDate) {
	return func(it *InputDate) {
		it.Value = v
	}
}

func WithInputDateErrorMessage(em string) func(*InputDate) {
	return func(it *InputDate) {
		it.ErrorMessage = em
	}
}

func WithInputDateRequired() func(*InputDate) {
	return func(it *InputDate) {
		it.IsRequired = true
	}
}

func WithInputDateLabel(l string) func(*InputDate) {
	return func(it *InputDate) {
		it.Label = l
	}
}

func WithInputDateLabelPosition(lp InputLabelPosition) func(*InputDate) {
	return func(it *InputDate) {
		it.LabelPosition = lp
	}
}

func WithInputDateLabelWidth[T LabelWidth](lw T) func(*InputDate) {
	return func(it *InputDate) {
		it.LabelWidth = lw
	}
}

func WithInputDateInputStyle(is InputStyle) func(*InputDate) {
	return func(it *InputDate) {
		it.InputStyle = is
	}
}

func WithInputDateHeight(beh BlockElementHeight) func(*InputDate) {
	return func(it *InputDate) {
		it.Height = beh
	}
}

func WithInputDateSeparator() func(*InputDate) {
	return func(it *InputDate) {
		it.Separator = true
	}
}

func WithInputDateSpacing(s Spacing) func(*InputDate) {
	return func(it *InputDate) {
		it.Spacing = s
	}
}

func WithInputDateNotVisible() func(*InputDate) {
	return func(it *InputDate) {
		it.IsVisible = false
	}
}
