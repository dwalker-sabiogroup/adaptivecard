package adaptivecard

type ActionToggleVisibility struct {
	Type           string          `json:"type"`
	TargetElements []TargetElement `json:"targetElements"`
	Title          string          `json:"title"`
	IconUrl        string          `json:"iconUrl"`
	ID             string          `json:"id"`
	Style          ActionStyle     `json:"style"`
	Tooltip        string          `json:"string"`
	IsEnabled      bool            `json:"isEnabled"`
	Mode           ActionMode      `json:"mode"`
}

func NewActionToggleVisibility(elements []TargetElement, opts ...func(*ActionToggleVisibility)) *ActionToggleVisibility {
	a := &ActionToggleVisibility{
		Type:           "Action.ToggleVisibility",
		TargetElements: elements,
		IsEnabled:      true,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

func WithActionToggleVisibilityTitle(t string) func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.Title = t
	}
}

func WithActionToggleVisibilityIconURL(u string) func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.IconUrl = u
	}
}

func WithActionToggleVisibilityID(i string) func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.ID = i
	}
}

func WithActionToggleVisibilityStyle(as ActionStyle) func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.Style = as
	}
}

func WithActionToggleVisibilityTooltip(t string) func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.Tooltip = t
	}
}

func WithActionToggleVisibilityDisabled() func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.IsEnabled = false
	}
}

func WithActionToggleVisibilityMode(am ActionMode) func(*ActionToggleVisibility) {
	return func(a *ActionToggleVisibility) {
		a.Mode = am
	}
}
