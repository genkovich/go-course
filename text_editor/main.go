package text_editor

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Text struct {
	originalText string
	rows         []string
}

func Start() {
	var userText string

	userText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
		"\nUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate " +
		"\nvelit esse cillum dolore eu fugiat nulla pariatur. " +
		"\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	preparedText := NewText(userText)

	emptyResult, _ := json.MarshalIndent(preparedText.rows, "", "  ")
	fmt.Println(string(emptyResult))

	substrings := preparedText.search("dolor")

	printSlice(substrings)
}

func NewText(userText string) Text {
	text := Text{originalText: userText}
	text.chunkOnRows()

	return text
}

func (t *Text) chunkOnRows() {
	t.rows = strings.Split(t.originalText, "\n")
}

func (t *Text) search(substr string) []string {
	var result []string
	for _, row := range t.rows {
		if strings.Contains(row, substr) {
			result = append(result, row)
		}
	}

	return result
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	for _, item := range s {
		fmt.Printf("Match found on: %s\n", item)
	}

}
