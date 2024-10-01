package ports

import (
	"github.com/jrabbit56/book-store/internal/core/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]domain.Book, error)
	GetByID(id uint) (*domain.Book, error)
	Create(book *domain.Book) error
	Update(book *domain.Book) error
	Delete(id uint) error
}

type UserRepository interface {
	CreateUser(user *domain.User) error
	LoginUser(user *domain.User) (string, error)
}

type OrderRepository interface {
	SaveOrder(order *domain.Order) error
	BeginTx() (*gorm.DB, error)
	UpdateBookStock(bookID uint, quantity int) error
	GetAllOrder() ([]domain.Order, error)
	GetOrderByID(id uint) (*domain.Order, error)
}

type InventoryRepository interface {
	CheckInventory(bookIDs []int) (map[int]int, error)
	UpdateInventory(bookID int, quantity int) error
	GetInventoryForBook(bookID int) (*domain.Inventory, error)
}
