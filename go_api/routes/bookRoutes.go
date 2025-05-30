package routes

import (
	"example.com/api/handlers"
	"example.com/api/middleware"
	"example.com/api/models"
	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")

	router.HandleFunc("/books",
		middleware.AuthMiddleware(
			middleware.RequireRole(models.RoleSeller)(
				handlers.CreateBook,
			),
		),
	).Methods("POST")

	router.HandleFunc("/books/{id}",
		middleware.AuthMiddleware(
			middleware.RequireRole(models.RoleSeller)(
				handlers.UpdateBook,
			),
		),
	).Methods("PUT")

	router.HandleFunc("/books/{id}",
		middleware.AuthMiddleware(
			middleware.RequireRole(models.RoleSeller)(
				handlers.DeleteBook,
			),
		),
	).Methods("DELETE")
}
