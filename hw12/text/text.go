package text

import (
	"course/hw12/text/processing"
)

type Text struct {
	content   string
	processor processing.TextProcessor
}

func CreateText(content string, processor processing.TextProcessor) *Text {
	return &Text{
		content:   content,
		processor: processor,
	}
}

func (t *Text) ChangeProcessor(tp processing.TextProcessor) {
	t.processor = tp
}

func (t *Text) Process() string {
	return t.processor.Process(t.content)
}
