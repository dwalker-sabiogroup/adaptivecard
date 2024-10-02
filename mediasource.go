package adaptivecard

type MediaSource struct {
	URL      string `json:"url"`
	MimeType string `json:"mimeType"`
}

func NewMediaSource(url string, opts ...func(*MediaSource)) *MediaSource {
	ms := &MediaSource{
		URL: url,
	}

	for _, o := range opts {
		o(ms)
	}

	return ms
}

func WithMediaSourceMimeType(mt string) func(*MediaSource) {
	return func(ms *MediaSource) {
		ms.MimeType = mt
	}
}
