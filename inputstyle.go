//go:generate go run github.com/dmarkham/enumer -type=InputStyle -json -transform=title-lower -trimprefix=InputStyle
package adaptivecard

type InputStyle int

const (
	InputStyleDefault InputStyle = iota
	InputStyleRevealOnHover
)
