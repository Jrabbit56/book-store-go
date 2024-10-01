package domain

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title        string  `json:"Title" example:"The Go Programming Language"`
	ISBN         string  `json:"ISBN" example:"978-0134190440"`
	AuthorID     int     `json:"AuthorID" example:"1"`
	TypeOfBookID int     `json:"TypeOfBookID" example:"2"`
	Price        float32 `json:"Price" example:"49.99"`
	Quantity     int     `json:"Quantity" example:"100"`
}

// SwaggerUser represents the user structure for Swagger documentation
type SwaggerBook struct {
	Title        string  `json:"Title" example:"The Go Programming Language"`
	ISBN         string  `json:"ISBN" example:"978-0134190440"`
	AuthorID     int     `json:"AuthorID" example:"1"`
	TypeOfBookID int     `json:"TypeOfBookID" example:"2"`
	Price        float32 `json:"Price" example:"49.99"`
	Quantity     int     `json:"Quantity" example:"100"`
}
