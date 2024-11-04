package adaptivecard

type Authentication struct {
	Text                  string                `json:"text"`
	ConnectionName        string                `json:"connectionName"`
	TokenExchangeResource TokenExchangeResource `json:"tokenExchangeResource"`
	Buttons               []AuthCardButton      `json:"buttons"`
}

func NewAuthentication(opts ...func(*Authentication)) *Authentication {
	r := &Authentication{
		Buttons: make([]AuthCardButton, 0),
	}

	for _, o := range opts {
		o(r)
	}

	return r
}

func WithAuthenticationText(t string) func(*Authentication) {
	return func(a *Authentication) {
		a.Text = t
	}
}

func WithAuthenticationConnectionName(cn string) func(*Authentication) {
	return func(a *Authentication) {
		a.ConnectionName = cn
	}
}

func WithAuthenticationTokenExchangeResource(ter TokenExchangeResource) func(*Authentication) {
	return func(a *Authentication) {
		a.TokenExchangeResource = ter
	}
}

func WithAuthenticationButtons(acb []AuthCardButton) func(*Authentication) {
	return func(a *Authentication) {
		a.Buttons = acb
	}
}
