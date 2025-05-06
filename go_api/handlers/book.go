package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/api/database"
	"example.com/api/models"
	"example.com/api/storage"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Book API")
}


func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range storage.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	result := database.DB.Create(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range storage.Books {
		if item.ID == params["id"] {
			json.NewDecoder(r.Body).Decode(&storage.Books[i])
			storage.Books[i].ID = params["id"] // ensure ID stays same
			json.NewEncoder(w).Encode(storage.Books[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range storage.Books {
		if item.ID == params["id"] {
			storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
