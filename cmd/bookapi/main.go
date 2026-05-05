package main

import (
	"log"
	"net/http"

	controller "github.com/srisubhash011/bookmanagement/controller"

	"github.com/gorilla/mux"
)

func main() {

	router := registerBookRoutes()

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}

func registerBookRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", controller.GetBookController).Methods("GET")
	router.HandleFunc("/book/search/{query}", controller.GetBookBySearchController).Methods("GET")
	router.HandleFunc("/book/add", controller.AddBookController).Methods("POST")
	return router
}
