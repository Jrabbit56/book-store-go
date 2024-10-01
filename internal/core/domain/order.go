package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PoNumber   string      `json:"PoNumber"`
	CustomerID int         `json:"CustomerID"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"Items"`
	Discount   float32     `json:"Discount"`
	TotalPrice float32     `json:"TotalPrice"`
	IsPayment  bool        `json:"IsPayment"`
}

// SwaggerUser represents the user structure for Swagger documentation
type SwaggerOrder struct {
	PoNumber   string             `json:"PoNumber"`
	CustomerID int                `json:"CustomerID"`
	Items      []SwaggerOrderItem `gorm:"foreignKey:OrderID" json:"Items"`
	Discount   float32            `json:"Discount"`
	TotalPrice float32            `json:"TotalPrice"`
	IsPayment  bool               `json:"IsPayment"`
}
