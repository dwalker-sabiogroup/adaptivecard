//go:generate go run github.com/dmarkham/enumer -type=TextInputStyle -json -transform=title-lower -trimprefix=TextInputStyle
package adaptivecard

type TextInputStyle int

const (
	TextInputStyleText TextInputStyle = iota
	TextInputStyleTel
	TextInputStyleUrl
	TextInputStyleEmail
	TextInputStylePassword
)
