package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"user-management-server/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	authService := services.AuthService{}
	user, err := authService.Authenticate(input.Email, input.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	generateService := services.JWTService{SecretKey: os.Getenv("JWT_SECRET_KEY")}
	token, err := generateService.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login success",
		"token":   token,
	})
}
