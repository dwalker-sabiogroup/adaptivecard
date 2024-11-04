package adaptivecard

type BackgroundImage struct {
	URL                 string              `json:"url"`
	FillMode            ImageFillMode       `json:"fillMode"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment"`
	VerticalAlignment   VerticalAlignment   `json:"verticalAlignment"`
}

func NewBackgroundImage(url string, opts ...func(*BackgroundImage)) *BackgroundImage {
	ics := &BackgroundImage{
		URL: url,
	}

	for _, o := range opts {
		o(ics)
	}

	return ics
}

func WithBackgroundImageFillMode(fm ImageFillMode) func(*BackgroundImage) {
	return func(bi *BackgroundImage) {
		bi.FillMode = fm
	}
}

func WithBackgroundImageVerticalAlignment(va VerticalAlignment) func(*BackgroundImage) {
	return func(bi *BackgroundImage) {
		bi.VerticalAlignment = va
	}
}

func WithBackgroundImageHorizontalAlignment(ha HorizontalAlignment) func(*BackgroundImage) {
	return func(bi *BackgroundImage) {
		bi.HorizontalAlignment = ha
	}
}
