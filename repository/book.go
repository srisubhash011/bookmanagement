package repository

import (
	"fmt"

	repositorymodel "github.com/srisubhash011/bookmanagement/models/repositorymodel"
)

var books = []repositorymodel.Book{
	{BookId: 1, BookName: "The Great Gatsby", BookDescription: "A novel by F. Scott Fitzgerald", BookStatus: 1},
	{BookId: 2, BookName: "To Kill a Mockingbird", BookDescription: "A novel by Harper Lee", BookStatus: 1},
	{BookId: 3, BookName: "1984", BookDescription: "A novel by George Orwell", BookStatus: 0},
}

func GetBook(bookId int) repositorymodel.Book {
	for _, book := range books {
		if book.BookId == bookId {

			return book
		}
	}

	return repositorymodel.Book{}

}

func ListAllBooks() []repositorymodel.Book {
	return books
}

func AddBook(book repositorymodel.Book) {

	books = append(books, book)
}

func UpdateBook(bookId int, book repositorymodel.Book) error {

	for i, b := range books {
		if b.BookId == bookId {
			books[i] = book
			return nil
		}
	}
	return fmt.Errorf("book with id %d not found", bookId)
}

func DeleteBook(bookId int) error {

	for i, book := range books {
		if book.BookId == bookId {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with id %d not found", bookId)
}
