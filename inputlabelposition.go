//go:generate go run github.com/dmarkham/enumer -type=InputLabelPosition -json -transform=title-lower -trimprefix=InputLabelPosition
package adaptivecard

type InputLabelPosition int

const (
	InputLabelPositionAbove InputLabelPosition = iota
	InputLabelPositionInline
)
