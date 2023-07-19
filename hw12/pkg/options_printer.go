package pkg

import (
	"github.com/manifoldco/promptui"
)

func PrintOptions(optionsTitle string, options map[string]string) (error, string) {
	kv := make([][]string, 0, len(options))
	for key, value := range options {
		kv = append(kv, []string{key, value})
	}

	values := make([]string, 0, len(options))
	for _, pair := range kv {
		values = append(values, pair[1])
	}

	prompt := promptui.Select{
		Label: optionsTitle,
		Items: values,
	}

	idx, _, err := prompt.Run()
	return err, kv[idx][0]
}
