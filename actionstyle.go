//go:generate go run github.com/dmarkham/enumer -type=ActionStyle -json -transform=title-lower -trimprefix=ActionStyle
package adaptivecard

type ActionStyle int

const (
	ActionStyleDefault ActionStyle = iota
	ActionStylePositive
	ActionStyleDestructive
)
