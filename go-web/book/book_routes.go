package book

import "github.com/gorilla/mux"

func SetRoutes(r *mux.Router) {
	s := r.PathPrefix("/books").Subrouter()
	s.HandleFunc("", addBook).Methods("POST")
	s.HandleFunc("", getBooks).Methods("GET")
	s.HandleFunc("/{isbn}", getBook).Methods("GET")
}
