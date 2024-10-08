package adaptivecard

type ImageSet struct {
	Type      string             `json:"type"`
	Images    []Image            `json:"images"`
	ImageSize ImageSize          `json:"imageSize"`
	Height    BlockElementHeight `json:"height"`
	Separator bool               `json:"separator"`
	Spacing   Spacing            `json:"spacing"`
	ID        string             `json:"id"`
	IsVisible bool               `json:"isVisible"`
}

func NewImageSet(images []Image, opts ...func(*ImageSet)) *ImageSet {
	ics := &ImageSet{
		Type:      "ImageSet",
		Images:    images,
		ImageSize: ImageSizeMedium,
		IsVisible: true,
	}

	for _, o := range opts {
		o(ics)
	}

	return ics
}

func WithImageSetSeparator() func(*ImageSet) {
	return func(i *ImageSet) {
		i.Separator = true
	}
}

func WithImageSetHeight(beh BlockElementHeight) func(*ImageSet) {
	return func(i *ImageSet) {
		i.Height = beh
	}
}

func WithImageSetSpacing(s Spacing) func(*ImageSet) {
	return func(i *ImageSet) {
		i.Spacing = s
	}
}

func WithImageSetID(id string) func(*ImageSet) {
	return func(i *ImageSet) {
		i.ID = id
	}
}

func WithImageSetNotVisible() func(*ImageSet) {
	return func(i *ImageSet) {
		i.IsVisible = false
	}
}
