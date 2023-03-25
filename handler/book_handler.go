package handler

import (
	"SimpleRestAPI/model"
	"errors"
	"github.com/google/uuid"
	"log"
)

type Bookhandler struct{}

var books []model.Book

func (bh Bookhandler) Add(bookReq model.BookRequest) (model.Book, error) {
	id := uuid.New()
	book := model.Book{BookId: id.String(), Title: bookReq.Title, Authors: bookReq.Authors}
	books = append(books, book)
	return book, nil
}

func (bh Bookhandler) List() []model.Book {
	return books
}

func (bh Bookhandler) Find(book_id string) (model.Book, error) {
	for _, book := range books {
		if book.BookId == book_id {
			log.Println("Found a matching book", book)
			return book, nil
		}
	}
	return model.Book{}, errors.New("Failed to fetch book with ID - " + book_id)
}

func (bh Bookhandler) Remove() {
	log.Println("TBD")
	return
}
