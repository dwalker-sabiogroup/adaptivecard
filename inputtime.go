package adaptivecard

type InputTime struct {
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

func NewInputTime(id string, opts ...func(*InputTime)) *InputTime {
	it := &InputTime{
		Type:      "Input.Time",
		ID:        id,
		IsVisible: true,
	}

	for _, o := range opts {
		o(it)
	}

	return it
}

func WithInputTimeMax(m string) func(*InputTime) {
	return func(it *InputTime) {
		it.Max = m
	}
}

func WithInputTimeMin(m string) func(*InputTime) {
	return func(it *InputTime) {
		it.Min = m
	}
}

func WithInputTimePlaceholder(p string) func(*InputTime) {
	return func(it *InputTime) {
		it.Placeholder = p
	}
}

func WithInputTimeValue(v string) func(*InputTime) {
	return func(it *InputTime) {
		it.Value = v
	}
}

func WithInputTimeErrorMessage(em string) func(*InputTime) {
	return func(it *InputTime) {
		it.ErrorMessage = em
	}
}

func WithInputTimeRequired() func(*InputTime) {
	return func(it *InputTime) {
		it.IsRequired = true
	}
}

func WithInputTimeLabel(l string) func(*InputTime) {
	return func(it *InputTime) {
		it.Label = l
	}
}

func WithInputTimeLabelPosition(lp InputLabelPosition) func(*InputTime) {
	return func(it *InputTime) {
		it.LabelPosition = lp
	}
}

func WithInputTimeLabelWidth[T LabelWidth](lw T) func(*InputTime) {
	return func(it *InputTime) {
		it.LabelWidth = lw
	}
}

func WithInputTimeInputStyle(is InputStyle) func(*InputTime) {
	return func(it *InputTime) {
		it.InputStyle = is
	}
}

func WithInputTimeHeight(beh BlockElementHeight) func(*InputTime) {
	return func(it *InputTime) {
		it.Height = beh
	}
}

func WithInputTimeSeparator() func(*InputTime) {
	return func(it *InputTime) {
		it.Separator = true
	}
}

func WithInputTimeSpacing(s Spacing) func(*InputTime) {
	return func(it *InputTime) {
		it.Spacing = s
	}
}

func WithInputTimeNotVisible() func(*InputTime) {
	return func(it *InputTime) {
		it.IsVisible = false
	}
}
