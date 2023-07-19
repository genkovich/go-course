package processing

import (
	"fmt"
	"strings"
)

type Counter interface {
	count(t string)
}

type Count struct{}

func (c *Count) Process(text string, targets []string) {

	var counter Counter = basicCounter{}

	for _, target := range targets {
		switch target {
		case "words":
			counter = CountWords{
				parent: counter,
			}
		case "unique":
			counter = UniqueWordsCounter{
				parent: counter,
			}
		case "whitespace":
			counter = WhiteSpacesCounter{
				parent: counter,
			}
		}
	}

	counter.count(text)
}

type basicCounter struct{}

func (p basicCounter) count(text string) {}

type CountWords struct {
	parent Counter
}

func (c CountWords) count(text string) {
	words := strings.Fields(text)
	fmt.Printf("Words: %d\n", len(words))
	c.parent.count(text)
}

type UniqueWordsCounter struct {
	parent Counter
}

func (c UniqueWordsCounter) count(text string) {
	words := strings.Fields(text)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}

	fmt.Printf("Unique words: %d\n", len(m))
	c.parent.count(text)
}

type WhiteSpacesCounter struct {
	parent Counter
}

func (c WhiteSpacesCounter) count(text string) {
	spaces := strings.Count(text, " ")
	fmt.Printf("White spaces: %d\n", spaces)
	c.parent.count(text)
}
