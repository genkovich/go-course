package book

import "fmt"

type Database struct {
	Books map[string]DatabaseItem
}

type DatabaseItem struct {
	Book
	ClientName string
	IsBorrowed bool
}

func NewDatabase(books []Book) *Database {
	bookMap := make(map[string]DatabaseItem)
	for _, book := range books {
		bookMap[book.Name] = DatabaseItem{Book: book}
	}

	return &Database{Books: bookMap}
}

func (d *Database) AddBook(book Book) {
	d.Books[book.Name] = DatabaseItem{Book: book}
}

func (d *Database) RemoveBook(bookName string) {
	delete(d.Books, bookName)
}

func (d *Database) MarkAsBorrowed(bookName string, clientName string) error {
	book, ok := d.Books[bookName]
	if !ok {
		return ErrBookNotFound
	}

	book.IsBorrowed = true
	book.ClientName = clientName
	d.Books[bookName] = book
	fmt.Println("Book", bookName, "is borrowed by", clientName)
	return nil
}

func (d *Database) MarkAsReturned(bookName string) error {
	book, ok := d.Books[bookName]
	if !ok {
		return ErrBookNotFound
	}

	book.IsBorrowed = false
	book.ClientName = ""
	d.Books[bookName] = book
	fmt.Println("Book", bookName, "is returned")
	return nil
}

func (d *Database) Search(bookName string) error {
	_, ok := d.Books[bookName]
	if !ok {
		return ErrBookNotFound
	}

	fmt.Println("Book", bookName, "is found")
	return nil
}

func (d *Database) IsBorrowed(bookName string) bool {
	book, ok := d.Books[bookName]
	if !ok {
		return false
	}

	return book.IsBorrowed
}

func (d *Database) IsBorrowedByClient(clientName string, bookName string) bool {
	book, ok := d.Books[bookName]
	if !ok {
		return false
	}

	return book.ClientName == clientName
}
