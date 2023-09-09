package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	UserID uint   `json:"user_id"`
}
