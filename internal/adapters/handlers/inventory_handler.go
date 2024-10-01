package handlers

import (
	"github.com/jrabbit56/book-store/internal/core/ports"
)

type InventoryHandler struct {
	service ports.InventoryService
}

func NewInventoryHandler(service ports.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}
