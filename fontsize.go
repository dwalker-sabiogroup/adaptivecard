//go:generate go run github.com/dmarkham/enumer -type=FontSize -json -transform=title-lower -trimprefix=FontSize
package adaptivecard

type FontSize int

const (
	FontSizeDefault FontSize = iota
	FontSizeSmall
	FontSizeMedium
	FontSizeLarge
	FontSizeExtraLarge
)
