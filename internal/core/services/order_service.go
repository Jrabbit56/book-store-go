package services

import (
	"fmt"

	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
	"gorm.io/gorm"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *domain.Order) error {
	// business logic core

	// Start a transaction
	tx, err := s.repo.BeginTx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var totalQuantity int
	var subtotal float32

	// Check stock and calculate totals
	for _, item := range order.Items {
		var book domain.Book
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&book, item.BookID).Error; err != nil {
			tx.Rollback()
			return err
		}
		if book.Quantity < item.Quantity {
			tx.Rollback()
			return fmt.Errorf("insufficient stock for book ID %d", item.BookID)
		}

		// Update stock
		if err := tx.Model(&domain.Book{}).Where("id = ?", item.BookID).
			Update("quantity", gorm.Expr("quantity - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return err
		}

		itemTotal := float32(item.Quantity) * item.Price
		subtotal += itemTotal
		totalQuantity += item.Quantity
	}

	// Apply initial discount
	discountedTotal := subtotal - order.Discount

	// Apply additional 3% discount if total quantity is more than 30
	if totalQuantity > 30 {
		additionalDiscount := discountedTotal * 0.03
		order.Discount += additionalDiscount
		discountedTotal -= additionalDiscount
	}

	// Set the final total price
	order.TotalPrice = discountedTotal

	fmt.Println(order.TotalPrice)

	// Save the order
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

func (s *OrderService) GetAllOrder() ([]domain.Order, error) {
	return s.repo.GetAllOrder()
}

func (s *OrderService) GetOrderWithItems(orderID uint) (*domain.Order, error) {
	return s.repo.GetOrderByID(orderID)
}
