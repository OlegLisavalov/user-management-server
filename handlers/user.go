package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"user-management-server/services"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		Password     string `json:"password"`
		ReferralCode string `json:"referral_code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := services.RegisterUser(input.Email, input.Name, input.Password, input.ReferralCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":       "User registered successfully",
		"user_id":       user.ID,
		"referral_code": user.ReferralCode,
	})
}

func GetUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := services.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"user_id":       user.ID,
		"email":         user.Email,
		"name":          user.Name,
		"referral_code": user.ReferralCode,
	})
}
