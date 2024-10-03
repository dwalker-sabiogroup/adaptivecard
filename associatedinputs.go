//go:generate go run github.com/dmarkham/enumer -type=AssociatedInputs -json -transform=title-lower -trimprefix=AssociatedInputs
package adaptivecard

type AssociatedInputs int

const (
	AssociatedInputsAuto AssociatedInputs = iota
	AssociatedInputsNone
)
