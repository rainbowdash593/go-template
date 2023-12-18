package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int `json:"id,primary_key" json:"id"`
	Name      string
	Email     string         `gorm:"uniqueIndex" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
