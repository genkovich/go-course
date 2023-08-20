package main

import (
	"course/hw21/library"
	book "course/hw21/library/book"
	"fmt"
)

func main() {
	books := []book.Book{
		{Id: "1", Name: "Little Prince"},
		{Id: "2", Name: "Harry Potter"},
		{Id: "3", Name: "Sherlock Holmes"},
		{Id: "4", Name: "The Lord of the Rings"},
		{Id: "5", Name: "The Hobbit"},
		{Id: "6", Name: "And Then There Were None"},
		{Id: "7", Name: "Dream of the Red Chamber"},
		{Id: "8", Name: "The Lion, the Witch and the Wardrobe"},
		{Id: "9", Name: "She: A History of Adventure"},
		{Id: "10", Name: "The Da Vinci Code"},
	}

	database := book.NewDatabase(books)
	shelf := book.NewShelf(books)
	librarian := library.Librarian{Name: "Steve"}
	lib := library.NewLibrary(&librarian, *shelf, *database)

	client := library.NewClient("John")
	fmt.Println("Books on the shelf:", lib.BookShelf.Books)

	err := lib.RequestBook("Little Prince", client)
	if err != nil {
		fmt.Println(err)
		return
	}

	// After a week...
	fmt.Println("One week later...")

	lib.ChangeLibrarian(&library.Librarian{Name: "Harry"})

	client.SwapBook()
	err = lib.ReturnBook(client)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Two weeks later...")
	client.SwapBook()
	err = lib.ReturnBook(client)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Books on the shelf:", lib.BookShelf.Books)
}
