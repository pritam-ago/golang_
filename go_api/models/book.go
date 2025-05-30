package models

type Book struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	UserID string `json:"user_id" gorm:"index"`
}
