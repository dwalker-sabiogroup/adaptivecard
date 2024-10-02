//go:generate go run github.com/dmarkham/enumer -type=ImageStyle -json -transform=title-lower -trimprefix=ImageStyle
package adaptivecard

type ImageStyle int

const (
	ImageStyleDefault ImageStyle = iota
	ImageStylePerson
)
