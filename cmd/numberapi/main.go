package main

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/srisubhash011/bookmanagement/services"
)

func main() {
	router := registerNumberRoutes()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func registerNumberRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/number/{n}", numberHandler).Methods(http.MethodGet)
	return router
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nStr := vars["n"]
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	primes := services.GetPrimeNumber(n)
	fibs := services.GetFibonacci(n)
	odds := services.GetOddNumber(n)

	unique := make(map[int]struct{})
	values := []int{}

	for _, v := range primes {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
			values = append(values, v)
		}
	}
	for _, v := range fibs {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
			values = append(values, v)
		}
	}
	for _, v := range odds {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
			values = append(values, v)
		}
	}

	sort.Ints(values)

	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = strconv.Itoa(v)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(strings.Join(strValues, " ")))
}
