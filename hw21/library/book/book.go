package book

import "fmt"

type Book struct {
	Id   string
	Name string
}

var ErrBookNotFound = fmt.Errorf("no book found in library database")
var ErrBookAlreadyBorrowed = fmt.Errorf("book is already borrowed")
var ErrBookIsLost = fmt.Errorf("book is lost")
var ErrDifferentClient = fmt.Errorf("book is borrowed by another client")

func (b *Book) String() string {
	return fmt.Sprintf("[%s]", b.Name)
}
