package main

import (
	"course/hw12/pkg"
	userText "course/hw12/text"
	"course/hw12/text/processing"
	"fmt"
	"os"
)

const (
	ThirdPartyPath   = "./hw12/third_party/"
	CountAllWords    = "count_all"
	CountUniqueWords = "count_unique"
)

func main() {
	var filename string

	fmt.Println("Enter filename")
	fmt.Scanln(&filename)

	fileContent, err := os.ReadFile(ThirdPartyPath + filename)
	if err != nil {
		fmt.Println("Error file reading")
		return
	}

	options := map[string]string{
		CountAllWords:    "Порахувати всі слова",
		CountUniqueWords: "Порахувати унікальні слова",
	}

	err, pickedStrategy := pkg.PrintOptions("Оберіть варіант", options)

	var strategy processing.TextProcessor

	switch pickedStrategy {
	case CountAllWords:
		strategy = &processing.CountWords{}
	case CountUniqueWords:
		strategy = &processing.CountUniqueWords{}
	default:
		strategy = &processing.CountWords{}
	}

	content := string(fileContent)

	text := userText.CreateText(content, strategy)
	fmt.Println(text.Process())

}
