package domain

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID  uint    `json:"OrderID"`
	BookID   int     `json:"BookID"`
	Quantity int     `json:"Quantity"`
	Price    float32 `json:"Price"`
}

// SwaggerUser represents the user structure for Swagger documentation
type SwaggerOrderItem struct {
	OrderID  uint    `json:"OrderID"`
	BookID   int     `json:"BookID"`
	Quantity int     `json:"Quantity"`
	Price    float32 `json:"Price"`
}
