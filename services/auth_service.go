package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"

	"user-management-server/database"
	"user-management-server/models"
)

type AuthService struct{}

func (as *AuthService) Authenticate(email, password string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

type JWTService struct {
	SecretKey string
}

func (js *JWTService) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.SecretKey))
}
