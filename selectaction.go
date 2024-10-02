//go:generate go run github.com/dmarkham/enumer -type=SelectAction -json -linecomment
package adaptivecard

type SelectAction int

const (
	SelectActionExecute          SelectAction = iota // Action.Execute
	SelectActionOpenURL                              // Action.OpenUrl
	SelectActionSubmit                               // Action.Submit
	SelectActionToggleVisibility                     //Action.ToggleVisibility
)
