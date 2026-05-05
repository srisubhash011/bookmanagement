package services

import (
	"fmt"
	"strings"

	repositorymodel "github.com/srisubhash011/bookmanagement/models/repositorymodel"
	servicesmodel "github.com/srisubhash011/bookmanagement/models/servicesmodel"
	"github.com/srisubhash011/bookmanagement/repository"
)

func GetBook(bookId int) servicesmodel.Book {
	var book = repository.GetBook(bookId)

	var serviceBook servicesmodel.Book
	serviceBook.BookId = book.BookId
	serviceBook.BookName = book.BookName
	serviceBook.BookDescription = book.BookDescription

	return serviceBook
}

func GetBookBySearch(search string) []servicesmodel.BookSearch {

	var books = repository.ListAllBooks()

	var searchResults []servicesmodel.BookSearch

	for _, book := range books {
		if strings.Contains(book.BookName, search) || strings.Contains(book.BookDescription, search) {
			fmt.Println("Book Name: ", book.BookName)
			var searchResult servicesmodel.BookSearch
			searchResult.BookName = book.BookName
			searchResults = append(searchResults, searchResult)
		}

	}

	return searchResults

}

func AddBook(book repositorymodel.Book) {
	repository.AddBook(book)
}
