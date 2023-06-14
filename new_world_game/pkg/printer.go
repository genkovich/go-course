package pkg

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"time"
)

const deleteRow = "\u001B[1A\u001B[2K"

func PrintText(text string) {
	rowCount := 0
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(10 * time.Millisecond)
		if char == '\r' {
			rowCount++

			fmt.Print("\nPress 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			fmt.Print(deleteRow)
			fmt.Print(deleteRow)
		}
	}

	fmt.Print("\nPress 'Enter' to finish...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	for i := 0; i < rowCount+2; i++ {
		fmt.Print(deleteRow)
	}
}

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
