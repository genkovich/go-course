package processing

type TextProcessor interface {
	Process(text string, targets []string)
}
