//go:generate go run github.com/dmarkham/enumer -type=ActionMode -json -transform=title-lower -trimprefix=ActionMode
package adaptivecard

type ActionMode int

const (
	ActionModePrimary ActionMode = iota
	ActionModeSecondary
)
