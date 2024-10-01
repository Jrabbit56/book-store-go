package domain

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Quantity int `json:"Quantity"` // Quantity in stock
}

type BookInventory struct {
	gorm.Model
	BookID      int `json:"BookID"`
	InventoryID uint
	Inventory   Inventory `gorm:"foreignKey:InventoryID"`
}

type InventoryStatus struct {
	BookID   int
	Quantity int
	InStock  bool
}
