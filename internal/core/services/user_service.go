package services

import (
	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) LoginUser(user *domain.User) (string, error) {
	return s.repo.LoginUser(user)
}
