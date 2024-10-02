package adaptivecard

type Media struct {
	Type           string             `json:"type"`
	Sources        []MediaSource      `json:"sources"`
	Poster         string             `json:"poster"`
	AltText        string             `json:"altText"`
	CaptionSources []CaptionSource    `json:"captionSources"`
	Height         BlockElementHeight `json:"height"`
	Separator      bool               `json:"separator"`
	Spacing        Spacing            `json:"spacing"`
	ID             string             `json:"id"`
	IsVisible      bool               `json:"isVisible"`
}

func NewMedia(sources []MediaSource, opts ...func(*Media)) *Media {
	m := &Media{
		Type:      "Media",
		Sources:   sources,
		IsVisible: true,
	}

	for _, o := range opts {
		o(m)
	}

	return m
}

func WithMediaPoster(p string) func(*Media) {
	return func(m *Media) {
		m.Poster = p
	}
}

func WithMediaAltText(at string) func(*Media) {
	return func(m *Media) {
		m.AltText = at
	}
}

func WithMediaCaptionSources(cs []CaptionSource) func(*Media) {
	return func(m *Media) {
		m.CaptionSources = cs
	}
}

func WithMediaHeight(bh BlockElementHeight) func(*Media) {
	return func(m *Media) {
		m.Height = bh
	}
}

func WithMediaSeparator() func(*Media) {
	return func(m *Media) {
		m.Separator = true
	}
}

func WithMediaSpacing(s Spacing) func(*Media) {
	return func(m *Media) {
		m.Spacing = s
	}
}

func WithMediaID(s string) func(*Media) {
	return func(m *Media) {
		m.ID = s
	}
}

func WithMediaNotVisible() func(*Media) {
	return func(m *Media) {
		m.IsVisible = false
	}
}
