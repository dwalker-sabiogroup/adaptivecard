//go:generate go run github.com/dmarkham/enumer -type=ImageSetStyle -json -transform=title-lower -trimprefix=ImageSetStyle
package adaptivecard

type ImageSetStyle int

const (
	ImageSetStyleDefault ImageSetStyle = iota
	ImageSetStyleStacked
	ImageSetStyleGrid
)
