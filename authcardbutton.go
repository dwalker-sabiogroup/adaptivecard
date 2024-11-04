package adaptivecard

type AuthCardButton struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Title string `json:"title"`
	Image string `json:"image"`
}

func NewAuthCardButton(t, value string, opts ...func(*AuthCardButton)) *AuthCardButton {
	acb := &AuthCardButton{
		Type:  t,
		Value: value,
	}

	for _, o := range opts {
		o(acb)
	}

	return acb
}

func WithAuthCardButtonTitle(t string) func(*AuthCardButton) {
	return func(acb *AuthCardButton) {
		acb.Title = t
	}
}

func WithAuthCardButtonImage(i string) func(*AuthCardButton) {
	return func(acb *AuthCardButton) {
		acb.Image = i
	}
}
