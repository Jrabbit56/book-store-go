package services

import (
	"errors"

	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
)

type InventoryService struct {
	repo ports.InventoryRepository
}

func NewInventoryService(repo ports.InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) CheckInventoryAvailability(bookIDs []int, requestedQuantities map[int]int) (map[int]bool, error) {
	inventoryMap, err := s.repo.CheckInventory(bookIDs)
	if err != nil {
		return nil, err
	}

	availability := make(map[int]bool)
	for bookID, requestedQty := range requestedQuantities {
		availableQty, exists := inventoryMap[bookID]
		availability[bookID] = exists && availableQty >= requestedQty
	}
	return availability, nil
}

func (s *InventoryService) UpdateInventory(bookID int, quantityChange int) error {
	inventory, err := s.repo.GetInventoryForBook(bookID)
	if err != nil {
		return err
	}

	newQuantity := inventory.Quantity + quantityChange
	if newQuantity < 0 {
		return errors.New("insufficient inventory")
	}

	return s.repo.UpdateInventory(bookID, newQuantity)
}

func (s *InventoryService) GetInventoryStatus(bookID int) (*domain.InventoryStatus, error) {
	inventory, err := s.repo.GetInventoryForBook(bookID)
	if err != nil {
		return nil, err
	}

	return &domain.InventoryStatus{
		BookID:   bookID,
		Quantity: inventory.Quantity,
		InStock:  inventory.Quantity > 0,
	}, nil
}
