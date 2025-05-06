package main

import (
	"log"
	"net/http"
	"example.com/api/routes"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
