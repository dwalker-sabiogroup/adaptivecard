package adaptivecard

type InputChoiceSet struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Choices       []InputChoice      `json:"choices"`
	ChoicesData   DataQuery          `json:"choices.data"`
	IsMultiSelect bool               `json:"isMultiSelect"`
	Style         ChoiceInputStyle   `json:"style"`
	Value         string             `json:"value"`
	Placeholder   string             `json:"placeholder"`
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

func NewInputChoiceSet(id string, opts ...func(*InputChoiceSet)) *InputChoiceSet {
	ics := &InputChoiceSet{
		Type:      "Input.ChoiceSet",
		ID:        id,
		IsVisible: true,
	}

	for _, o := range opts {
		o(ics)
	}

	return ics
}

func WithInputChoiceSetChoices(c []InputChoice) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Choices = c
	}
}

func WithInputChoiceSetDataQuery(dq DataQuery) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.ChoicesData = dq
	}
}

func WithInputChoiceSetMultiSelect() func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.IsMultiSelect = true
	}
}

func WithInputChoiceSetStyle(s ChoiceInputStyle) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Style = s
	}
}

func WithInputChoiceSetWrap() func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Wrap = true
	}
}

func WithInputChoiceSetPlaceholder(p string) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Placeholder = p
	}
}

func WithInputChoiceSetValue(v string) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Value = v
	}
}

func WithInputChoiceSetErrorMessage(em string) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.ErrorMessage = em
	}
}

func WithInputChoiceSetRequired() func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.IsRequired = true
	}
}

func WithInputChoiceSetLabel(l string) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Label = l
	}
}

func WithInputChoiceSetLabelPosition(lp InputLabelPosition) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.LabelPosition = lp
	}
}

func WithInputChoiceSetLabelWidth[T LabelWidth](lw T) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.LabelWidth = lw
	}
}

func WithInputChoiceSetInputStyle(is InputStyle) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.InputStyle = is
	}
}

func WithInputChoiceSetHeight(beh BlockElementHeight) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Height = beh
	}
}

func WithInputChoiceSetSeparator() func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Separator = true
	}
}

func WithInputChoiceSetSpacing(s Spacing) func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.Spacing = s
	}
}

func WithInputChoiceSetNotVisible() func(*InputChoiceSet) {
	return func(ics *InputChoiceSet) {
		ics.IsVisible = false
	}
}
