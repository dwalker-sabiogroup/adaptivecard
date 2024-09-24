//go:generate go run github.com/dmarkham/enumer -type=Spacing -json -transform=title-lower -trimprefix=Spacing
package adaptivecard

type Spacing int

const (
	SpacingDefault Spacing = iota
	SpacingNone
	SpacingSmall
	SpacingMedium
	SpacingLarge
	SpacingExtraLarge
	SpacingPadding
)
