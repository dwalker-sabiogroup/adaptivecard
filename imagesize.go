//go:generate go run github.com/dmarkham/enumer -type=ImageSize -json -transform=title-lower -trimprefix=ImageSize
package adaptivecard

type ImageSize int

const (
	ImageSizeAuto ImageSize = iota
	ImageSizeStretch
	ImageSizeSmall
	ImageSizeMedium
	ImageSizeLarge
)
