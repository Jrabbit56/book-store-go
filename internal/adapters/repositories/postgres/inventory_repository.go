package postgres

import (
	"github.com/jrabbit56/book-store/internal/core/domain"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) CheckInventory(bookIDs []int) (map[int]int, error) {
	var bookInventories []domain.BookInventory
	result := r.db.Where("book_id IN ?", bookIDs).Preload("Inventory").Find(&bookInventories)
	if result.Error != nil {
		return nil, result.Error
	}

	inventoryMap := make(map[int]int)
	for _, bi := range bookInventories {
		inventoryMap[bi.BookID] = bi.Inventory.Quantity
	}
	return inventoryMap, nil
}

func (r *InventoryRepository) UpdateInventory(bookID int, quantity int) error {
	var bookInventory domain.BookInventory
	result := r.db.Where("book_id = ?", bookID).First(&bookInventory)
	if result.Error != nil {
		return result.Error
	}

	return r.db.Model(&domain.Inventory{}).Where("id = ?", bookInventory.InventoryID).
		Update("quantity", quantity).Error
}

func (r *InventoryRepository) GetInventoryForBook(bookID int) (*domain.Inventory, error) {
	var bookInventory domain.BookInventory
	result := r.db.Where("book_id = ?", bookID).Preload("Inventory").First(&bookInventory)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bookInventory.Inventory, nil
}
