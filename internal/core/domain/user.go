package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"` // Ignore Gorm fields in Swagger
	Email      string                 `gorm:"unique"`
	Password   string                 `json:"Password"`
	Role       int                    `json:"Role"`
}

// SwaggerUser represents the user structure for Swagger documentation
type SwaggerUser struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
	Role     int    `json:"role" example:"1"`
}
