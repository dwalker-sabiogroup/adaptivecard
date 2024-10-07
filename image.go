package adaptivecard

type ImageHeight interface {
	string | BlockElementHeight
}

type Image struct {
	Type                string              `json:"type"`
	URL                 string              `json:"url"`
	AltText             string              `json:"altText"`
	BackgroundColor     string              `json:"backgroundColor"`
	Height              interface{}         `json:"height"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment"`
	SelectAction        interface{}         `json:"selectAction"`
	Size                ImageSize           `json:"size"`
	Style               ImageStyle          `json:"style"`
	Width               string              `json:"width"`
	Separator           bool                `json:"separator"`
	Spacing             Spacing             `json:"spacing"`
	ID                  string              `json:"id"`
	IsVisible           bool                `json:"isVisible"`
}

func NewImage(url string, opts ...func(*Image)) *Image {
	i := &Image{
		Type:      "Image",
		URL:       url,
		Height:    BlockElementHeightAuto,
		IsVisible: true,
	}

	for _, o := range opts {
		o(i)
	}

	return i
}

func WithImageAltText(at string) func(*Image) {
	return func(i *Image) {
		i.AltText = at
	}
}

func WithImageBackgroupColor(bc string) func(*Image) {
	return func(i *Image) {
		i.BackgroundColor = bc
	}
}

func WithImageHeight[T ImageHeight](h T) func(*Image) {
	return func(i *Image) {
		i.Height = h
	}
}

func WithImageHorizontalAlignment(ha HorizontalAlignment) func(*Image) {
	return func(i *Image) {
		i.HorizontalAlignment = ha
	}
}

func WithImageSelectAction[T SelectAction](sa T) func(*Image) {
	return func(i *Image) {
		i.SelectAction = sa
	}
}

func WithImageSize(s ImageSize) func(*Image) {
	return func(i *Image) {
		i.Size = s
	}
}

func WithImageStyle(s ImageStyle) func(*Image) {
	return func(i *Image) {
		i.Style = s
	}
}

func WithImageWidth(w string) func(*Image) {
	return func(i *Image) {
		i.Width = w
	}
}

func WithImageSeparator() func(*Image) {
	return func(i *Image) {
		i.Separator = true
	}
}

func WithImageSpacing(s Spacing) func(*Image) {
	return func(i *Image) {
		i.Spacing = s
	}
}

func WithImageID(id string) func(*Image) {
	return func(i *Image) {
		i.ID = id
	}
}

func WithImageNotVisible() func(*Image) {
	return func(i *Image) {
		i.IsVisible = false
	}
}
