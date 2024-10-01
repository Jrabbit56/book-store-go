package services

import (
	"errors"
	"testing"
	"time"

	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

// Mock implementation of BookRepository
type mockBookRepo struct {
	CreateFunc  func(book *domain.Book) error
	GetAllFunc  func() ([]domain.Book, error)
	GetByIDFunc func(id uint) (*domain.Book, error)
	UpdateFunc  func(book *domain.Book) error
	DeleteFunc  func(id uint) error
}

func (m *mockBookRepo) Create(book *domain.Book) error {
	return m.CreateFunc(book)
}

func (m *mockBookRepo) GetAll() ([]domain.Book, error) {
	return m.GetAllFunc()
}

func (m *mockBookRepo) GetByID(id uint) (*domain.Book, error) {
	return m.GetByIDFunc(id)
}

func (m *mockBookRepo) Update(book *domain.Book) error {
	return m.UpdateFunc(book)
}

func (m *mockBookRepo) Delete(id uint) error {
	return m.DeleteFunc(id)
}

func TestBookService(t *testing.T) {

	t.Run("GetAllBooks", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			expectedBooks := []domain.Book{
				{
					Model:        gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
					Title:        "Book 1",
					ISBN:         "111-1111111111",
					AuthorID:     1,
					TypeOfBookID: 1,
					Price:        29.99,
					Quantity:     50,
				},
				{
					Model:        gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
					Title:        "Book 2",
					ISBN:         "222-2222222222",
					AuthorID:     2,
					TypeOfBookID: 2,
					Price:        39.99,
					Quantity:     75,
				},
			}
			repo := &mockBookRepo{
				GetAllFunc: func() ([]domain.Book, error) {
					return expectedBooks, nil
				},
			}
			service := NewBookService(repo)
			books, err := service.GetAllBooks()
			assert.Nil(t, err)
			assert.Equal(t, expectedBooks, books)
		})

		t.Run("Failure", func(t *testing.T) {
			repo := &mockBookRepo{
				GetAllFunc: func() ([]domain.Book, error) {
					return nil, errors.New("failed to get books")
				},
			}
			service := NewBookService(repo)
			books, err := service.GetAllBooks()
			assert.NotNil(t, err)
			assert.Nil(t, books)
			assert.Equal(t, "failed to get books", err.Error())
		})
	})

	t.Run("GetBookByID", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			expectedBook := &domain.Book{
				Model:        gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Title:        "Found Book",
				ISBN:         "333-3333333333",
				AuthorID:     3,
				TypeOfBookID: 3,
				Price:        59.99,
				Quantity:     25,
			}
			repo := &mockBookRepo{
				GetByIDFunc: func(id uint) (*domain.Book, error) {
					return expectedBook, nil
				},
			}
			service := NewBookService(repo)
			book, err := service.GetBook(1)
			assert.Nil(t, err)
			assert.Equal(t, expectedBook, book)
		})

		t.Run("Failure", func(t *testing.T) {
			repo := &mockBookRepo{
				GetByIDFunc: func(id uint) (*domain.Book, error) {
					return nil, errors.New("book not found")
				},
			}
			service := NewBookService(repo)
			book, err := service.GetBook(999)
			assert.NotNil(t, err)
			assert.Nil(t, book)
			assert.Equal(t, "book not found", err.Error())
		})
	})

	t.Run("UpdateBook", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			updatedBook := &domain.Book{
				Model:        gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Title:        "Updated Book Title",
				ISBN:         "444-4444444444",
				AuthorID:     4,
				TypeOfBookID: 4,
				Price:        69.99,
				Quantity:     150,
			}
			repo := &mockBookRepo{
				UpdateFunc: func(book *domain.Book) error {
					// Simulate updating the book
					book.UpdatedAt = time.Now()
					return nil
				},
			}
			service := NewBookService(repo)
			err := service.UpdateBook(updatedBook)
			assert.Nil(t, err)
			assert.Equal(t, "Updated Book Title", updatedBook.Title)
			assert.Equal(t, "444-4444444444", updatedBook.ISBN)
		})

		t.Run("Failure", func(t *testing.T) {
			repo := &mockBookRepo{
				UpdateFunc: func(book *domain.Book) error {
					return errors.New("failed to update book")
				},
			}
			service := NewBookService(repo)
			book := &domain.Book{Model: gorm.Model{ID: 1}}
			err := service.UpdateBook(book)
			assert.NotNil(t, err)
			assert.Equal(t, "failed to update book", err.Error())
		})
	})

	t.Run("DeleteBook", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			repo := &mockBookRepo{
				DeleteFunc: func(id uint) error {
					return nil
				},
			}
			service := NewBookService(repo)
			err := service.DeleteBook(1)
			assert.Nil(t, err)
		})

		t.Run("Failure", func(t *testing.T) {
			repo := &mockBookRepo{
				DeleteFunc: func(id uint) error {
					return errors.New("failed to delete book")
				},
			}
			service := NewBookService(repo)
			err := service.DeleteBook(999)
			assert.NotNil(t, err)
			assert.Equal(t, "failed to delete book", err.Error())
		})
	})
}
