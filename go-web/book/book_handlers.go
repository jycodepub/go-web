package book

import (
	"encoding/json"
	"go-web/internal/app/web"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var bookRepository *BookRepositry

func init() {
	uri, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		uri = "mongodb://localhost:27017"
	}
	log.Printf("Loaded Mongodb URI: %s\n", uri)
	bookRepository = NewBookRepository(uri)
}

func getBooks(w http.ResponseWriter, q *http.Request) {
	books, err := bookRepository.GetAll()
	if err != nil {
		web.TextResponse(w, 500, "Database Errors")
	} else {
		web.JsonResponse(w, 200, books)
	}
}

func getBook(w http.ResponseWriter, q *http.Request) {
	isbn := mux.Vars(q)["isbn"]
	book, err := bookRepository.Get(isbn)
	if err != nil {
		web.TextResponse(w, 500, "Database Error")
	} else {
		web.JsonResponse(w, 200, book)
	}
}

func addBook(w http.ResponseWriter, q *http.Request) {
	body, err := io.ReadAll(q.Body)
	if err != nil {
		web.TextResponse(w, 500, "IO Error")
		return
	}
	defer q.Body.Close()

	var book Book
	if err := json.Unmarshal(body, &book); err != nil {
		web.TextResponse(w, 500, "Json Error: "+err.Error()+"\nbody: "+string(body))
		return
	}

	id, err := bookRepository.Add(book)
	if err != nil {
		web.TextResponse(w, 500, "Database Error")
	} else {
		web.TextResponse(w, 200, "Added book: "+id)
	}
}
