package todo

import "github.com/gorilla/mux"

func SetRoutes(r *mux.Router) {
	r.HandleFunc("/todo", todoPageHandler).Methods("GET")
	r.HandleFunc("/todo", addTask).Methods("POST")
	r.HandleFunc("/todo/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/todo/{id}", updateTask).Methods("PUT")
}
