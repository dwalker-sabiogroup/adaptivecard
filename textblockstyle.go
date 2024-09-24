//go:generate go run github.com/dmarkham/enumer -type=TextBlockStyle -json -transform=title-lower -trimprefix=TextBlockStyle
package adaptivecard

type TextBlockStyle int

const (
	TextBlockStyleDefault TextBlockStyle = iota
	TextBlockStyleHeading
)
