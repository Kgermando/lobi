package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
 
	Fullname  string    `gorm:"not null" json:"fullname"`
	Email     string    `gorm:"not null" json:"email"` 
	Message   string    `gorm:"not null" json:"message"`
	IsRead    bool      `gorm:"not null" json:"is_read"` 
}
