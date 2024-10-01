package handlers

import (
	"bytes"
	"encoding/json"

	// "errors"
	// "io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBookService struct {
	mock.Mock
}

func (m *mockBookService) GetAllBooks() ([]domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m *mockBookService) GetBook(id uint) (*domain.Book, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *mockBookService) CreateBook(book *domain.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *mockBookService) UpdateBook(book *domain.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *mockBookService) DeleteBook(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestBookHandler(t *testing.T) {

	t.Run("CreateBook", func(t *testing.T) {
		mockService := new(mockBookService)
		handler := NewBookHandler(mockService)
		app := fiber.New()
		app.Post("/books", handler.CreateBook)

		t.Run("Success", func(t *testing.T) {
			book := &domain.Book{
				Title:        "The Go Programming Language",
				ISBN:         "978-0134190440",
				AuthorID:     1,
				TypeOfBookID: 2,
				Price:        49.99,
				Quantity:     100,
			}

			mockService.On("CreateBook", mock.AnythingOfType("*domain.Book")).Return(nil)

			body, _ := json.Marshal(book)
			req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

			mockService.AssertExpectations(t)
		})

	})

	t.Run("GetAllBooks", func(t *testing.T) {
		mockService := new(mockBookService)
		handler := NewBookHandler(mockService)
		app := fiber.New()
		app.Get("/books", handler.GetAllBooks)

		t.Run("Success", func(t *testing.T) {
			books := []domain.Book{
				{
					Title:        "The Go Programming Language",
					ISBN:         "978-0134190440",
					AuthorID:     1,
					TypeOfBookID: 2,
					Price:        49.99,
					Quantity:     100,
				},
				{
					Title:        "Another Go Book",
					ISBN:         "978-0321563842",
					AuthorID:     2,
					TypeOfBookID: 3,
					Price:        39.99,
					Quantity:     50,
				},
			}
			mockService.On("GetAllBooks").Return(books, nil)

			req := httptest.NewRequest("GET", "/books", nil)
			resp, _ := app.Test(req)

			assert.Equal(t, fiber.StatusOK, resp.StatusCode)

			var responseBooks []domain.Book
			err := json.NewDecoder(resp.Body).Decode(&responseBooks)
			assert.Nil(t, err)                    // Ensure no error occurred while decoding
			assert.Equal(t, books, responseBooks) // Compare expected books with response

			mockService.AssertExpectations(t)
		})

	})

}
