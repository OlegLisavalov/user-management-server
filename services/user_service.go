package services

import (
	"errors"
	"github.com/google/uuid"
	"user-management-server/database"
	"user-management-server/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, name, password, referralCode string) (*models.User, error) {
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("User with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Error hashing password")
	}

	newReferralCode, err := GenerateUniqueReferralCode()
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		ID:           uuid.New().String(),
		Email:        email,
		Name:         name,
		Password:     string(hashedPassword),
		ReferralCode: newReferralCode,
	}

	if referralCode != "" {
		var referrer models.User
		if err := database.DB.Where("referral_code = ?", referralCode).First(&referrer).Error; err == nil {
			referrer.Points += 5
			newUser.Points += 7
			if err := database.DB.Save(&referrer).Error; err != nil {
				return nil, errors.New("Failed to update referrer points")
			}
			newUser.ReferrerID = &referrer.ID
		}
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, errors.New("Failed to register user")
	}

	return &newUser, nil
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetTopUsersByPoints(limit int) ([]models.User, error) {
	var users []models.User
	err := database.DB.Order("points desc").Limit(limit).Find(&users).Error
	return users, err
}
