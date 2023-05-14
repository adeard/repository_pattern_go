package book

import "time"

type Book struct {
	ID          int
	Title       string `json:"title" `
	Description string `json:"description" `
	Price       int    `json:"price" `
	Rating      int    `json:"rating"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
