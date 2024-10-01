package postgres

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/jrabbit56/book-store/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Replace the plaintext password with the hashed password
	user.Password = string(hashedPassword)
	// Use r.db to create the user in the database
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) LoginUser(user *domain.User) (string, error) {
	// Fetch the user by email
	selectedUser := new(domain.User)

	result := r.db.Where("email =?", user.Email).First(selectedUser)
	if result.Error != nil {
		// Return a more specific error if the user is not found
		return "", result.Error
	}
	// Compare the password hashes
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		// If password comparison fails, return an error
		return "", errors.New("invalid password")
	}

	load := godotenv.Load()
	if load != nil {
		log.Fatal("Error loading .env file")
	}
	// Generate a JWT token with the user's ID and expiration time
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["role"] = selectedUser.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expiration time
	// Sign the token with the secret key
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	// Return the generated JWT token
	return t, nil
}
