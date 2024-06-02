package models

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

// User struct
type User struct {
	gorm.Model

	Email         string   `gorm:"uniqueIndex;not null" json:"email"`
	Password      string   `gorm:"not null" json:"password"`
	Fullname      string   `json:"fullname"`
	Address       string   `json:"address"`
	Telephone     string   `json:"telephone"`
	EmailVerified bool     `json:"email_verified"`
	Role          string   `json:"role"`
	IsActive      bool     `json:"is_active"` 
}
