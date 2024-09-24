//go:generate go run github.com/dmarkham/enumer -type=FontType -json -transform=title-lower -trimprefix=FontType
package adaptivecard

type FontType int

const (
	FontTypeDefault FontType = iota
	FontTypeMonospace
)
