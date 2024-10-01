package domain

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	WriterName string `json:"WriterName"`
}
