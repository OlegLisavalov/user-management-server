package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string  `json:"id" gorm:"primaryKey"` // ID будет строкой (например, UUID)
	Name         string  `json:"name"`
	Email        string  `json:"email" gorm:"unique"`
	Password     string  `json:"password"`
	Points       int     `json:"points"`
	ReferralCode string  `json:"referral_code"`
	ReferrerID   *string `json:"referrer_id"` // ReferrerID тоже строка (например, UUID)
}

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Points    int    `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	UserID    uint
}
