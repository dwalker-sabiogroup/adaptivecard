package adaptivecard

type TargetElement struct {
	ElementID string `json:"elementID"`
	IsVisible bool   `json:"isVisible"`
}

func NewTargetElement(id string, opts ...func(*TargetElement)) *TargetElement {
	i := &TargetElement{
		ElementID: id,
	}

	for _, o := range opts {
		o(i)
	}

	return i
}

func WithTargetElementNotVisible() func(*TargetElement) {
	return func(te *TargetElement) {
		te.IsVisible = false
	}
}

func WithTargetElementVisible() func(*TargetElement) {
	return func(te *TargetElement) {
		te.IsVisible = true
	}
}
