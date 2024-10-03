package adaptivecard

type ActionExecute struct {
	Type             string                 `json:"type"`
	Verb             string                 `json:"verb"`
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

func NewActionExecute(opts ...func(*ActionExecute)) *ActionExecute {
	a := &ActionExecute{
		Data:      make(map[string]interface{}),
		Type:      "Action.Execute",
		IsEnabled: true,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

func WithActionExecuteVerb(v string) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.Verb = v
	}
}

func WithActionExecuteData(d map[string]interface{}) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.Data = d
	}
}

func WithActionExecuteAssociatedInputs(ai AssociatedInputs) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.AssociatedInputs = ai
	}
}

func WithActionExecuteTitle(t string) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.Title = t
	}
}

func WithActionExecuteIconURL(u string) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.IconUrl = u
	}
}

func WithActionExecuteID(i string) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.ID = i
	}
}

func WithActionExecuteStyle(as ActionStyle) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.Style = as
	}
}

func WithActionExecuteTooltip(t string) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.Tooltip = t
	}
}

func WithActionExecuteDisabled() func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.IsEnabled = false
	}
}

func WithActionExecuteMode(am ActionMode) func(*ActionExecute) {
	return func(a *ActionExecute) {
		a.Mode = am
	}
}
