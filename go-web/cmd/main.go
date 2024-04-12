package main

import (
	"go-web/book"
	"go-web/internal/app"
	"go-web/internal/app/config"
	"go-web/todo"
	"go-web/web/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Set up routes
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	router.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	book.SetRoutes(router)
	todo.SetRoutes(router)

	// Static files
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Start the webapp
	webapp := app.NewApp(router)
	cfg := config.GetConfig()
	webapp.Start(cfg)
}
