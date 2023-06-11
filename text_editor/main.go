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
	preparedText.chunkOnRows()

	emptyResult, _ := json.MarshalIndent(preparedText.rows, "", "  ")
	fmt.Println(string(emptyResult))

	preparedText.search("dolor")

}

func NewText(userText string) Text {
	text := Text{originalText: userText}

	return text
}

func (t *Text) chunkOnRows() {
	t.rows = strings.Split(t.originalText, "\n")
}

func (t *Text) search(substr string) {
	for i, row := range t.rows {
		if strings.Contains(row, substr) {
			fmt.Printf("Match found on row %d: %s\n", i+1, row)
		}
	}
}
