package adaptivecard

type ActionShowCard struct {
	Type      string       `json:"type"`
	Card      AdaptiveCard `json:"card"`
	Title     string       `json:"title"`
	IconUrl   string       `json:"iconUrl"`
	ID        string       `json:"id"`
	Style     ActionStyle  `json:"style"`
	Tooltip   string       `json:"string"`
	IsEnabled bool         `json:"isEnabled"`
	Mode      ActionMode   `json:"mode"`
}

func NewActionShowCard(opts ...func(*ActionShowCard)) *ActionShowCard {
	a := &ActionShowCard{
		Type:      "Action.ShowCard",
		IsEnabled: true,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

func WithActionShowCardCard(ac AdaptiveCard) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.Card = ac
	}
}

func WithActionShowCardTitle(t string) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.Title = t
	}
}

func WithActionShowCardIconURL(u string) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.IconUrl = u
	}
}

func WithActionShowCardID(i string) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.ID = i
	}
}

func WithActionShowCardStyle(as ActionStyle) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.Style = as
	}
}

func WithActionShowCardTooltip(t string) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.Tooltip = t
	}
}

func WithActionShowCardDisabled() func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.IsEnabled = false
	}
}

func WithActionShowCardMode(am ActionMode) func(*ActionShowCard) {
	return func(a *ActionShowCard) {
		a.Mode = am
	}
}
