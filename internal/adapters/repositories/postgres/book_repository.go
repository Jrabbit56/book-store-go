package postgres

import (
	"errors"

	"github.com/jrabbit56/book-store/internal/core/domain"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *BookRepository) GetByID(id uint) (*domain.Book, error) {
	var book domain.Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *BookRepository) Create(book *domain.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) Update(book *domain.Book) error {
	result := r.db.Model(book).Where("id = ?", book.ID).Updates(book)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("book was updated by another user, please retry")
	}

	return nil
}

func (r *BookRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Book{}, id).Error
}
