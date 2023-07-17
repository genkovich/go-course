package text

import (
	"course/hw12/text/processing"
)

type Text struct {
	content   string
	processor processing.TextProcessor
	targets   []string
}

func CreateText(content string, processor processing.TextProcessor) *Text {
	return &Text{
		content:   content,
		processor: processor,
		targets:   []string{},
	}
}

func (t *Text) AddTarget(strategy string) {
	t.targets = append(t.targets, strategy)
}

func (t *Text) ChangeProcessor(tp processing.TextProcessor) {
	t.processor = tp
}

func (t *Text) Process() {
	t.processor.Process(t.content, t.targets)
}
