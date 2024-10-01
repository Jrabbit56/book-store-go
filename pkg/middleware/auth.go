package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Role int

const (
	Admin Role = iota + 1
	User
)

func AuthRequired(allowedRoles ...Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")

		if err := godotenv.Load(); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error loading .env file")
		}

		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

		token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
		}

		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
		}

		userRole, ok := (*claims)["role"].(float64)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid user role")
		}

		for _, role := range allowedRoles {
			if Role(userRole) == role {
				return c.Next()
			}
		}

		return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
	}
}