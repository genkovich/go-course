package processing

import (
	"strconv"
	"strings"
)

type CountUniqueWords struct{}

func (c *CountUniqueWords) Process(text string) string {
	words := strings.Fields(text)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return strconv.Itoa(len(m))
}
