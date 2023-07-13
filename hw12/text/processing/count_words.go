package processing

import (
	"strconv"
	"strings"
)

type CountWords struct{}

func (c *CountWords) Process(text string) string {
	words := strings.Fields(text)
	countOfWords := len(words)

	return strconv.Itoa(countOfWords)
}
