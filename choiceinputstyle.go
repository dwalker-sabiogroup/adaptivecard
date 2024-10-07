//go:generate go run github.com/dmarkham/enumer -type=ChoiceInputStyle -json -transform=title-lower -trimprefix=ChoiceInputStyle
package adaptivecard

type ChoiceInputStyle int

const (
	ChoiceInputStyleCompact ChoiceInputStyle = iota
	ChoiceInputStyleExpanded
	ChoiceInputStyleFiltered
)
