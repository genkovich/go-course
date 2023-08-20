package library

import (
	"course/hw21/library/book"
	"fmt"
)

type Library struct {
	Librarian    *Librarian
	BookShelf    book.Shelf
	BookDatabase book.Database
	Visitors     []Client
}

func NewLibrary(librarian *Librarian, bookShelf book.Shelf, bookDatabase book.Database) *Library {
	return &Library{Librarian: librarian, BookShelf: bookShelf, BookDatabase: bookDatabase}
}

func (l *Library) ChangeLibrarian(librarian *Librarian) {
	fmt.Println("Librarian changed from", l.Librarian.Name, "to", librarian.Name)
	l.Librarian = librarian
}

func (l *Library) RequestBook(bookName string, client *Client) error {
	err := l.BookDatabase.Search(bookName)
	if err != nil {
		return err
	}

	if l.BookDatabase.IsBorrowed(bookName) {
		return book.ErrBookAlreadyBorrowed
	}

	pickedBook, err := l.BookShelf.PickBook(bookName)
	if err != nil {
		return err
	}

	err = l.BookDatabase.MarkAsBorrowed(bookName, client.Name)
	if err != nil {
		return err
	}

	err = client.Borrow(pickedBook)
	if err != nil {
		_ = l.BookDatabase.MarkAsReturned(bookName)
		return err
	}

	fmt.Println("Book", bookName, "is borrowed by", client.Name)
	return nil
}

func (l *Library) ReturnBook(client *Client) error {
	borrowedBook := client.BorrowedBook()

	err := l.BookDatabase.Search(borrowedBook.Name)
	if err != nil {
		return err
	}

	ok := l.BookDatabase.IsBorrowedByClient(client.Name, borrowedBook.Name)
	if !ok {
		return book.ErrDifferentClient
	}

	err = client.Return(borrowedBook)
	if err != nil {
		return err
	}

	l.BookShelf.AddBook(borrowedBook)
	err = l.BookDatabase.MarkAsReturned(borrowedBook.Name)
	if err != nil {
		return err
	}

	return nil
}
