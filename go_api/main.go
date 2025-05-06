package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"example.com/api/database"
	"example.com/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	router := routes.SetupRoutes()
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
