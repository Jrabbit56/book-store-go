package domain

import (
	"gorm.io/gorm"
)

type TypeOfBook struct {
	gorm.Model
	TypeOfBooK string `json:"TypeOfBooK"`
}
