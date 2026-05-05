package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/srisubhash011/bookmanagement/models/repositorymodel"
	"github.com/srisubhash011/bookmanagement/services"
)

func GetBookController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	bookid, _ := strconv.Atoi(id)
	book := services.GetBook(bookid)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBookBySearchController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := params["query"]
	books := services.GetBookBySearch(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBookController(w http.ResponseWriter, r *http.Request) {
	var book repositorymodel.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	services.AddBook(book)
	w.WriteHeader(http.StatusCreated)
}

/*
func UpdateBookController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	bookid, _ := strconv.Atoi(id)
	var book repositorymodel.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	services.UpdateBook(bookid, book)
	w.WriteHeader(http.StatusNoContent)
}
*/
