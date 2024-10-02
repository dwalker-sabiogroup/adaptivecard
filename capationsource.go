package adaptivecard

type CaptionSource struct {
	Label    string `json:"label"`
	MimeType string `json:"mimeType"`
	URL      string `json:"url"`
}

func NewCaptionSource(mimeType, url, label string) *CaptionSource {
	return &CaptionSource{
		URL:      url,
		MimeType: mimeType,
		Label:    label,
	}
}
