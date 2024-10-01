package postgres

import (
	"errors"
	"testing"

	"github.com/jrabbit56/book-store/internal/core/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBookRepository_Create(t *testing.T) {
	// Mock database setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}
	repo := NewBookRepository(gormDB)

	// Success case
	t.Run("success", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "books"`).WillReturnRows(sqlmock.NewRows(
			[]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := repo.Create(&domain.Book{
			Title:        "The Python Programming Language",
			ISBN:         "978-0134190440",
			AuthorID:     1,
			TypeOfBookID: 2,
			Price:        49.99,
			Quantity:     100,
		})
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Failure case
	t.Run("failure", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectQuery("INSERT INTO").WillReturnError(errors.New("database errors"))
		mock.ExpectRollback()

		err := repo.Create(&domain.Book{
			Title:        "The Python Programming Language",
			ISBN:         "978-0134190440",
			AuthorID:     1,
			TypeOfBookID: 2,
			Price:        49.99,
			Quantity:     100,
		})
		assert.Error(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
