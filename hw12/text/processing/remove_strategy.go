package processing

import (
	"fmt"
	"strings"
	"unicode"
)

type Remover interface {
	remove(t string) string
}

type Remove struct{}

func (c *Remove) Process(text string, targerts []string) {
	var remover Remover = basicRemover{}

	for _, target := range targerts {
		switch target {
		case "words":
			remover = RemoveWords{
				parent: remover,
			}
		case "unique":
			remover = RemoveUniqueWords{
				parent: remover,
			}
		case "whitespace":
			remover = RemoveWhitespaces{
				parent: remover,
			}
		}
	}

	fmt.Println(remover.remove(text))

}

type basicRemover struct{}

func (p basicRemover) remove(text string) string {
	return text
}

type RemoveWords struct {
	parent Remover
}

func (c RemoveWords) remove(text string) string {
	words := strings.Fields(text)
	for _, word := range words {
		isWord := true
		for _, letter := range word {
			if !unicode.IsLetter(letter) {
				isWord = false
			}
		}
		if isWord {
			text = strings.ReplaceAll(text, word, "")
		}
	}

	return c.parent.remove(text)
}

type RemoveUniqueWords struct {
	parent Remover
}

func (c RemoveUniqueWords) remove(text string) string {
	words := strings.Fields(text)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}

	var result []string
	for _, word := range words {
		if m[word] != 1 {
			result = append(result, word)
		}
	}

	text = strings.Join(result, " ")
	return c.parent.remove(text)

}

type RemoveWhitespaces struct {
	parent Remover
}

func (c RemoveWhitespaces) remove(text string) string {
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	return c.parent.remove(text)
}
