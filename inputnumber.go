package adaptivecard

type InputNumber struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Max           int                `json:"max"`
	Min           int                `json:"min"`
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

func NewInputNumber(id string, opts ...func(*InputNumber)) *InputNumber {
	it := &InputNumber{
		Type:      "Input.Number",
		ID:        id,
		IsVisible: true,
	}

	for _, o := range opts {
		o(it)
	}

	return it
}

func WithInputNumberMax(m int) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Max = m
	}
}

func WithInputNumberMin(m int) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Min = m
	}
}

func WithInputNumberPlaceholder(p string) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Placeholder = p
	}
}

func WithInputNumberValue(v string) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Value = v
	}
}

func WithInputNumberErrorMessage(em string) func(*InputNumber) {
	return func(it *InputNumber) {
		it.ErrorMessage = em
	}
}

func WithInputNumberRequired() func(*InputNumber) {
	return func(it *InputNumber) {
		it.IsRequired = true
	}
}

func WithInputNumberLabel(l string) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Label = l
	}
}

func WithInputNumberLabelPosition(lp InputLabelPosition) func(*InputNumber) {
	return func(it *InputNumber) {
		it.LabelPosition = lp
	}
}

func WithInputNumberLabelWidth[T LabelWidth](lw T) func(*InputNumber) {
	return func(it *InputNumber) {
		it.LabelWidth = lw
	}
}

func WithInputNumberInputStyle(is InputStyle) func(*InputNumber) {
	return func(it *InputNumber) {
		it.InputStyle = is
	}
}

func WithInputNumberHeight(beh BlockElementHeight) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Height = beh
	}
}

func WithInputNumberSeparator() func(*InputNumber) {
	return func(it *InputNumber) {
		it.Separator = true
	}
}

func WithInputNumberSpacing(s Spacing) func(*InputNumber) {
	return func(it *InputNumber) {
		it.Spacing = s
	}
}

func WithInputNumberNotVisible() func(*InputNumber) {
	return func(it *InputNumber) {
		it.IsVisible = false
	}
}
