package main

import (
	_ "fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/jrabbit56/book-store/docs"
	"github.com/jrabbit56/book-store/internal/adapters/handlers"
	"github.com/jrabbit56/book-store/internal/adapters/repositories/postgres"
	_ "github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/services"
	"github.com/jrabbit56/book-store/pkg/database"
)

func authRequired(allowedRoles ...int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")

		load := godotenv.Load()
		if load != nil {
			log.Fatal("Error loading .env file")
		}
		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

		token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		userRole, ok := (*claims)["role"].(float64)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		for _, role := range allowedRoles {
			if int(userRole) == role {
				return c.Next()
			}
		}

		return c.SendStatus(fiber.StatusForbidden)
	}
}

// @title			Book store API
// @version		1.0
// @description	This is a sample Book store for Fiber
// @termsOfService	http://swagger.io/terms/
func main() {

	db := database.SetupDatabase()
	// Set up the core service and adapters

	orderRepo := postgres.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	userRepo := postgres.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	bookRepo := postgres.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// api := app.Group("/api")
	//Users routes

	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.LoginUser)

	//Manage bookshelf routes
	api := app.Group("/api")
	api.Use("/books", authRequired(1))
	api.Get("/books", bookHandler.GetAllBooks)
	api.Get("/books/:id", bookHandler.GetBook)
	api.Post("/books", bookHandler.CreateBook)
	api.Put("/books/:id", bookHandler.UpdateBook)
	api.Delete("/books/:id", bookHandler.DeleteBook)

	//Manage ordering of books routes
	app.Use("/orders", authRequired(1, 2))
	api.Post("/orders", orderHandler.CreateOrder)
	api.Get("/orders", orderHandler.GetAllOrder)
	api.Get("/orders/:id", orderHandler.GetOrder)

	log.Fatal(app.Listen(":8080"))
}
