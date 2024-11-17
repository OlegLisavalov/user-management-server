package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string  `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name"`
	Email        string  `json:"email" gorm:"unique"`
	Password     string  `json:"password"`
	Points       int     `json:"points"`
	ReferralCode string  `json:"referral_code" gorm:"unique"`
	ReferrerID   *string `json:"referrer_id"`
}

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Points    int    `json:"points"`
	UserID    string `json:"user_id"`
	Completed bool   `json:"completed"`
}

type LeaderboardUser struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Points int    `json:"points"`
}
