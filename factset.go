package adaptivecard

type FactSet struct {
	Type      string             `json:"type"`
	Facts     []Fact             `json:"facts"`
	Height    BlockElementHeight `json:"height"`
	Separator bool               `json:"separator"`
	Spacing   Spacing            `json:"spacing"`
	ID        string             `json:"id"`
	IsVisible bool               `json:"isVisible"`
}

func NewFactSet(facts []Fact, opts ...func(*FactSet)) *FactSet {
	ics := &FactSet{
		Type:      "FactSet",
		Facts:     facts,
		IsVisible: true,
	}

	for _, o := range opts {
		o(ics)
	}

	return ics
}

func WithFactSetSeparator() func(*FactSet) {
	return func(i *FactSet) {
		i.Separator = true
	}
}

func WithFactSetHeight(beh BlockElementHeight) func(*FactSet) {
	return func(i *FactSet) {
		i.Height = beh
	}
}

func WithFactSetSpacing(s Spacing) func(*FactSet) {
	return func(i *FactSet) {
		i.Spacing = s
	}
}

func WithFactSetID(id string) func(*FactSet) {
	return func(i *FactSet) {
		i.ID = id
	}
}

func WithFactSetNotVisible() func(*FactSet) {
	return func(i *FactSet) {
		i.IsVisible = false
	}
}
