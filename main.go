package main

import (
	"course/hw12/pkg"
	userText "course/hw12/text"
	"course/hw12/text/processing"
	"fmt"
	"os"
)

const (
	ThirdPartyPath = "./hw12/third_party/"
	Count          = "count"
	Remove         = "remove"
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
		Count:  "Порахувати",
		Remove: "Видалити",
	}

	err, pickedStrategy := pkg.PrintOptions("Оберіть варіант", options)

	var strategy processing.TextProcessor

	switch pickedStrategy {
	case Count:
		strategy = &processing.Count{}
	case Remove:
		strategy = &processing.Remove{}
	default:
		strategy = &processing.Count{}
	}

	content := string(fileContent)

	targets := map[string]string{
		"words":      "Слова",
		"unique":     "Унікальні слова",
		"whitespace": "Пробіли",
		"end":        "Завершити",
	}

	text := userText.CreateText(content, strategy)

	for {
		err, pickedTarget := pkg.PrintOptions("Оберіть варіант", targets)
		if err != nil {
			fmt.Println(err)
			return
		}
		if pickedTarget == "end" {
			break
		}

		text.AddTarget(pickedTarget)

		delete(targets, pickedTarget)
	}

	text.Process()
}
