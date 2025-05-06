package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"example.com/api/models"
	"example.com/api/storage"
)
	
// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Books)
}

// Get a single book
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

// Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	storage.Books = append(storage.Books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
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

// Delete a book
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
