package postgres

import (
	"errors"

	"github.com/jrabbit56/book-store/internal/core/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) BeginTx() (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (r *OrderRepository) SaveOrder(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) UpdateBookStock(bookID uint, quantity int) error {
	result := r.db.Model(&domain.Book{}).Where("id = ?", bookID).UpdateColumn("quantity", gorm.Expr("quantity - ?", quantity))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("book not found or insufficient stock")
	}
	return nil
}

func (r *OrderRepository) GetAllOrder() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
