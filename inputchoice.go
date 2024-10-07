package adaptivecard

type InputChoice struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

func NewInputChoice(title, value string) *InputChoice {
	return &InputChoice{
		Title: title,
		Value: value,
	}
}
