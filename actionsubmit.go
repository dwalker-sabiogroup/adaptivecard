package adaptivecard

type ActionSubmit struct {
	Type             string                 `json:"type"`
	Data             map[string]interface{} `json:"data"`
	AssociatedInputs AssociatedInputs       `json:"associatedInputs"`
	Title            string                 `json:"title"`
	IconUrl          string                 `json:"iconUrl"`
	ID               string                 `json:"id"`
	Style            ActionStyle            `json:"style"`
	Tooltip          string                 `json:"string"`
	IsEnabled        bool                   `json:"isEnabled"`
	Mode             ActionMode             `json:"mode"`
}

func NewActionSubmit(opts ...func(*ActionSubmit)) *ActionSubmit {
	a := &ActionSubmit{
		Data:      make(map[string]interface{}),
		Type:      "Action.Submit",
		IsEnabled: true,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

func WithActionSubmitData(d map[string]interface{}) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.Data = d
	}
}

func WithActionSubmitAssociatedInputs(ai AssociatedInputs) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.AssociatedInputs = ai
	}
}

func WithActionSubmitTitle(t string) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.Title = t
	}
}

func WithActionSubmitIconURL(u string) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.IconUrl = u
	}
}

func WithActionSubmitID(i string) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.ID = i
	}
}

func WithActionSubmitStyle(as ActionStyle) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.Style = as
	}
}

func WithActionSubmitTooltip(t string) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.Tooltip = t
	}
}

func WithActionSubmitDisabled() func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.IsEnabled = false
	}
}

func WithActionSubmitMode(am ActionMode) func(*ActionSubmit) {
	return func(a *ActionSubmit) {
		a.Mode = am
	}
}
