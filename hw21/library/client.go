package library

import (
	"course/hw21/library/book"
	"fmt"
)

type Client struct {
	Name         string
	borrowedBook *book.Book
	myBooks      map[string]*book.Book
}

func NewClient(name string) *Client {
	myBooks := make(map[string]*book.Book)
	myBooks["book1"] = &book.Book{Name: "book1"}

	return &Client{
		Name:    name,
		myBooks: myBooks,
	}
}

func (c *Client) Borrow(b *book.Book) error {
	if c.borrowedBook != nil {
		return fmt.Errorf("client %s already have a book %s", c.Name, c.borrowedBook.Name)
	}

	c.borrowedBook = b
	c.myBooks[b.Name] = b
	fmt.Printf("client %s borrowed book %s\n", c.Name, b.Name)
	return nil
}

func (c *Client) BorrowedBook() *book.Book {
	return c.borrowedBook
}

func (c *Client) SwapBook() {
	var bookNames []string
	for name := range c.myBooks {
		bookNames = append(bookNames, name)
	}

	var currentIndex int
	for i, name := range bookNames {
		if name == c.borrowedBook.Name {
			currentIndex = i
			break
		}
	}

	nextIndex := (currentIndex + 1) % len(bookNames)
	nextBookName := bookNames[nextIndex]

	c.borrowedBook = c.myBooks[nextBookName]
	fmt.Printf("client %s swapped to book %s\n", c.Name, c.borrowedBook.Name)
}

func (c *Client) Return(b *book.Book) error {
	if c.borrowedBook == nil {
		return fmt.Errorf("client %s has no book", c.Name)
	}

	if c.borrowedBook.Name != b.Name {
		return fmt.Errorf("client %s has another book %s", c.Name, c.borrowedBook.Name)
	}

	c.borrowedBook = nil
	fmt.Printf("client %s returned book %s\n", c.Name, b.Name)
	return nil
}
