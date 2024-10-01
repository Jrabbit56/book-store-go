package services

import (
	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
)

type BookService struct {
	repo ports.BookRepository
}

func NewBookService(repo ports.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]domain.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBook(id uint) (*domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) CreateBook(book *domain.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) UpdateBook(book *domain.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
