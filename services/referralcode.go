package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"user-management-server/database"
	"user-management-server/models"
)

func GenerateUniqueReferralCode() (string, error) {
	for {
		code := uuid.New().String()[:8]
		var user models.User
		if err := database.DB.Where("referral_code = ?", code).First(&user).Error; err == gorm.ErrRecordNotFound {
			return code, nil
		}
	}
	return "", errors.New("Failed to generate unique referral code")
}
