package adaptivecard

type Action interface {
	ActionExecute | ActionOpenURL | ActionSubmit | ActionToggleVisibility | ActionShowCard
}

type ActionSet struct {
	Type      string             `json:"type"`
	Actions   []interface{}      `json:"actions"`
	Height    BlockElementHeight `json:"height"`
	Separator bool               `json:"separator"`
	Spacing   Spacing            `json:"spacing"`
	ID        string             `json:"id"`
	IsVisible bool               `json:"isVisible"`
}

func NewActionSet(opts ...func(*ActionSet)) *ActionSet {
	ics := &ActionSet{
		Type:      "ActionSet",
		IsVisible: true,
	}

	for _, o := range opts {
		o(ics)
	}

	return ics
}

func WithActionSetAction[T Action](a T) func(*ActionSet) {
	return func(i *ActionSet) {
		i.Actions = append(i.Actions, a)
	}
}

func WithActionSetSeparator() func(*ActionSet) {
	return func(i *ActionSet) {
		i.Separator = true
	}
}

func WithActionSetHeight(beh BlockElementHeight) func(*ActionSet) {
	return func(i *ActionSet) {
		i.Height = beh
	}
}

func WithActionSetSpacing(s Spacing) func(*ActionSet) {
	return func(i *ActionSet) {
		i.Spacing = s
	}
}

func WithActionSetID(id string) func(*ActionSet) {
	return func(i *ActionSet) {
		i.ID = id
	}
}

func WithActionSetNotVisible() func(*ActionSet) {
	return func(i *ActionSet) {
		i.IsVisible = false
	}
}
