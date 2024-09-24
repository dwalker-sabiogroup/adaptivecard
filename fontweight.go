//go:generate go run github.com/dmarkham/enumer -type=FontWeight -json -transform=title-lower -trimprefix=FontWeight
package adaptivecard

type FontWeight int

const (
	FontWeightDefault FontWeight = iota
	FontWeightLighter
	FontWeightBolder
)
