package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./cw12/text.txt")
	if err != nil {
		panic(err)
	}

	var prnt ContentPrinter = basicContentPrinter{}

	var answer string

	fmt.Println("add lower?")
	fmt.Scanln(&answer)

	if answer == "y" {
		prnt = lowercaseContentPrinter{
			parent: prnt,
		}
	}
	prnt.printFileContent(string(content))

}

type ContentPrinter interface {
	printFileContent(c string)
}

type newLineContentPrinter struct {
	parent ContentPrinter
}

func (p newLineContentPrinter) printFileContent(c string) {
	p.parent.printFileContent(c + "\n")
}

type noSpacesContentPrinter struct {
	parent ContentPrinter
}

func (p noSpacesContentPrinter) printFileContent(c string) {
	p.parent.printFileContent(strings.TrimSpace(c))
}

type lowercaseContentPrinter struct {
	parent ContentPrinter
}

func (p lowercaseContentPrinter) printFileContent(c string) {
	p.parent.printFileContent(strings.ToLower(c))
}

type basicContentPrinter struct{}

func (p basicContentPrinter) printFileContent(c string) {
	fmt.Println(c)
}
