package adaptivecard

type Refresh struct {
	Action  ActionExecute `json:"action"`
	Expires string        `json:"expires"`
	UserIDs []string      `json:"userIds"`
}

func NewRefresh(opts ...func(*Refresh)) *Refresh {
	r := &Refresh{
		UserIDs: make([]string, 0),
	}

	for _, o := range opts {
		o(r)
	}

	return r
}

func WithRefreshAction(ae ActionExecute) func(*Refresh) {
	return func(r *Refresh) {
		r.Action = ae
	}
}

func WithRefreshExpires(e string) func(*Refresh) {
	return func(r *Refresh) {
		r.Expires = e
	}
}

func WithRefreshUserIDs(u []string) func(*Refresh) {
	return func(r *Refresh) {
		r.UserIDs = u
	}
}
