package adaptivecard

type Metadata struct {
	WebURL string `json:"webUrl"`
}

func NewMetadata(opts ...func(*Metadata)) *Metadata {
	m := &Metadata{}

	for _, o := range opts {
		o(m)
	}

	return m
}

func WithMetadataWebURL(wu string) func(*Metadata) {
	return func(r *Metadata) {
		r.WebURL = wu
	}
}
