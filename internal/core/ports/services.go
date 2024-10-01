package ports

import "github.com/jrabbit56/book-store/internal/core/domain"

type BookService interface {
	GetAllBooks() ([]domain.Book, error)
	GetBook(id uint) (*domain.Book, error)
	CreateBook(book *domain.Book) error
	UpdateBook(book *domain.Book) error
	DeleteBook(id uint) error
}

type UserService interface {
	RegisterUser(user *domain.User) error
	LoginUser(user *domain.User) (string, error)
}

type OrderService interface {
	CreateOrder(order *domain.Order) error
	GetAllOrder() ([]domain.Order, error)
	GetOrderWithItems(id uint) (*domain.Order, error)
}

type InventoryService interface {
	CheckInventoryAvailability(bookIDs []int, requestedQuantities map[int]int) (map[int]bool, error)
	UpdateInventory(bookID int, quantityChange int) error
	GetInventoryStatus(bookID int) (*domain.InventoryStatus, error)
}
