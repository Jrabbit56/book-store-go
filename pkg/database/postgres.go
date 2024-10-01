package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv"
	"github.com/jrabbit56/book-store/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// for Production
// const (
// 	db   = "db"
// 	port = 5432
// )

// For Development
const (
	host     = "localhost"   // or the Docker service name if running in another container
	port     = 5432          // default PostgreSQL port
	user     = "mydev"       // as defined in docker-compose.yml
	password = "mypassword"  // as defined in docker-compose.yml
	dbname   = "devdatabase" // as defined in docker-compose.yml
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {

	// //Load .env file (Production)
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// // Read database configuration from environment variables (Production)
	// user := os.Getenv("POSTGRES_USER")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// dbname := os.Getenv("POSTGRES_DB")

	// // Configure your PostgreSQL database details here (Production)
	// dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=5432 sslmode=disable",
	// 	db, user, password, dbname)

	//For Development
	dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=5432 sslmode=disable",
		host, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&domain.User{}, &domain.Book{}, &domain.Author{}, &domain.Order{}, &domain.TypeOfBook{}, &domain.OrderItem{})

	return db

}
