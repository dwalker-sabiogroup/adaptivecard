//go:generate go run github.com/dmarkham/enumer -type=HorizontalAlignment -json -transform=title-lower -trimprefix=HorizontalAlignment
package adaptivecard

type HorizontalAlignment int

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)
