//go:generate go run github.com/dmarkham/enumer -type=Color -json -transform=title-lower -trimprefix=Color
package adaptivecard

type Color int

const (
	ColorDefault Color = iota
	ColorDark
	ColorLight
	ColorAccent
	ColorGood
	ColorWarning
	ColorAttention
)
