package main

import (
	"fmt"
	"os"
	"strconv"

	//repositorymodel "github.com/srisubhash011/bookmanagement/models/repositorymodel"
	"github.com/srisubhash011/bookmanagement/models/servicesmodel"
	"github.com/srisubhash011/bookmanagement/services"
)

func main() {

	var args = os.Args

	fmt.Println(args[1])
	fmt.Println(args[2])

	if args[1] == "+book" {
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid book ID")
			return
		}
		var book = getBook(id)
		fmt.Println("Book ID:", book.BookId)
		fmt.Println("Book Name:", book.BookName)
		fmt.Println("Book Description:", book.BookDescription)
	}

	/*
		fmt.Println("Welcome to Book Management System")

		fmt.Println("\nGetting book with ID 1:")
		var book = services.GetBook(1)
		fmt.Println("Book ID:", book.BookId)
		fmt.Println("Book Name:", book.BookName)
		fmt.Println("Book Description:", book.BookDescription)

		fmt.Println("\n Search result for 'novel':")
		var searchResults = services.GetBookBySearch("nove")

		for _, result := range searchResults {
			fmt.Println("Book Name:", result.BookName)
		}

		fmt.Println("\nAdding a new book:")
		go services.AddBook(repositorymodel.Book{BookId: 4, BookName: "The Catcher in the Rye", BookDescription: "A novel by J.D. Salinger"})

		fmt.Println("\nGetting book with ID 4:")
		book = services.GetBook(4)

		fmt.Println("Book ID:", book.BookId)
		fmt.Println("Book Name:", book.BookName)
		fmt.Println("Book Description:", book.BookDescription)
	*/

}

func getBook(id int) servicesmodel.Book {
	var book = services.GetBook(id)
	return book
}
