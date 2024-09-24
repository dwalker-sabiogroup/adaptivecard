package adaptivecard

type ImageHeight interface {
	string | BlockElementHeight
}

type Image struct {
	Type            string      `json:"type"`
	URL             string      `json:"url"`
	AltText         string      `json:"altText"`
	BackgroundColor string      `json:"backgroundColor"`
	Height          interface{} `json:"height"`
}

func NewImage(url string, opts ...func(*Image)) *Image {
	i := &Image{
		Type: "Image",
		URL:  url,
	}

	for _, o := range opts {
		o(i)
	}

	return i
}

func WithHeight[T ImageHeight](h T) func(*Image) {
	return func(i *Image) {
		i.Height = h
	}
}
