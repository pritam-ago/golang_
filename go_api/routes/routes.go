package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	
)

func SetupRoutes() http.Handler {
	router := mux.NewRouter()
	
	HomeRoutes(router)
	BookRoutes(router)

	return router
}
