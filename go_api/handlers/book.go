package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/api/database"
	"example.com/api/middleware"
	"example.com/api/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)
	result := database.DB.Where("id = ?", params["id"]).First(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Author == "" || book.Year == 0 {
		http.Error(w, "Title, Author and Year are required", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	book.ID = uuid.New().String()
	book.UserID = userID

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
	var book models.Book
	params := mux.Vars(r)
	result := database.DB.Where("id = ?", params["id"]).First(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Verify ownership
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	if book.UserID != userID {
		http.Error(w, "Not authorized to update this book", http.StatusForbidden)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result = database.DB.Save(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)
	result := database.DB.Where("id = ?", params["id"]).First(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Verify ownership
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	if book.UserID != userID {
		http.Error(w, "Not authorized to delete this book", http.StatusForbidden)
		return
	}

	result = database.DB.Delete(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}
