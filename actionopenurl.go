package adaptivecard

type ActionOpenURL struct {
	Type      string      `json:"type"`
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	IconUrl   string      `json:"iconUrl"`
	ID        string      `json:"id"`
	Style     ActionStyle `json:"style"`
	Tooltip   string      `json:"string"`
	IsEnabled bool        `json:"isEnabled"`
	Mode      ActionMode  `json:"mode"`
}

func NewActionOpenURL(url string, opts ...func(*ActionOpenURL)) *ActionOpenURL {
	a := &ActionOpenURL{
		Type:      "Action.OpenUrl",
		URL:       url,
		IsEnabled: true,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

func WithActionOpenURLTitle(t string) func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.Title = t
	}
}

func WithActionOpenURLIconURL(u string) func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.IconUrl = u
	}
}

func WithActionOpenURLID(i string) func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.ID = i
	}
}

func WithActionOpenURLStyle(as ActionStyle) func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.Style = as
	}
}

func WithActionOpenURLTooltip(t string) func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.Tooltip = t
	}
}

func WithActionOpenURLDisabled() func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.IsEnabled = false
	}
}

func WithActionOpenURLMode(am ActionMode) func(*ActionOpenURL) {
	return func(a *ActionOpenURL) {
		a.Mode = am
	}
}
