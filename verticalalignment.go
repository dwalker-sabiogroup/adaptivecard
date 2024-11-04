//go:generate go run github.com/dmarkham/enumer -type=VerticalAlignment -json -transform=title-lower -trimprefix=VerticalAlignment
package adaptivecard

type VerticalAlignment int

const (
	VerticalAlignmentTop VerticalAlignment = iota
	VerticalAlignmentCenter
	VerticalAlignmentBottom
)
