package models

import "gorm.io/gorm"

type NewsLetter struct {
	gorm.Model

	Email string `gorm:"not null" json:"email"`
}