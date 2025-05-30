package routes

import (
	"example.com/api/handlers"
	"github.com/gorilla/mux"
)

func HomeRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.Home).Methods("GET")
}

