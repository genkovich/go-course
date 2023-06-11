package main

import "fmt"

type Book struct {
	title         string
	author        string
	numberOfPages int
}

type BookShelf struct {
	FirstBook Book
}

type ProgrammingBook struct {
	Book
	ProgrammingLanguage string
}

func (b *Book) AddPage() {
	b.numberOfPages++
}

func (b *Book) GetPages() int {
	return b.numberOfPages
}

func (b *Book) GetTitle() string {
	return b.title
}

func main() {
	book := Book{
		title:         "Go Lang",
		author:        "SomeAuthor",
		numberOfPages: 200,
	}

	book.author = "asdsad"
	fmt.Println(book)
}
