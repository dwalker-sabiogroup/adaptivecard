//go:generate go run github.com/dmarkham/enumer -type=ImageFillMode -json -transform=title-lower -trimprefix=ImageFillMode
package adaptivecard

type ImageFillMode int

const (
	ImageFillModeCover ImageFillMode = iota
	ImageFillModeRepeatHorizontally
	ImageFillModeRepeatVertically
	ImageFillModeRepeat
)
