//go:generate go run github.com/dmarkham/enumer -type=BlockElementHeight -json -transform=title-lower -trimprefix=BlockElementHeight
package adaptivecard

type BlockElementHeight int

const (
	BlockElementHeightAuto BlockElementHeight = iota
	BlockElementHeightStretch
)
