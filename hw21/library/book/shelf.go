package book

import "fmt"

type Shelf struct {
	Books []*Book
}

func NewShelf(books []Book) *Shelf {
	bookPointers := make([]*Book, len(books))
	for bookIndex := range books {
		bookPointers[bookIndex] = &books[bookIndex]
	}

	return &Shelf{Books: bookPointers}
}

func (s *Shelf) AddBook(book *Book) {
	s.Books = append(s.Books, book)
	fmt.Println("Book", book.Name, "is added to shelf")
}

func (s *Shelf) PickBook(bookName string) (*Book, error) {
	var pickedBook *Book

	for bookIndex := range s.Books {
		if s.Books[bookIndex].Name == bookName {
			pickedBook = s.Books[bookIndex]
			s.Books = append(s.Books[:bookIndex], s.Books[bookIndex+1:]...)
			break
		}
	}

	if pickedBook == nil {
		return nil, ErrBookIsLost
	}

	fmt.Println("Book", bookName, "is picked from shelf")
	return pickedBook, nil
}
