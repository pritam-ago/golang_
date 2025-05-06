package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"example.com/api/handlers"
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	return router
}
