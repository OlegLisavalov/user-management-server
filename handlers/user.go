package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management-server/database"
	"user-management-server/models"
	"user-management-server/services"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
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

func GetUserStatus(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := services.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func CheckReferrer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	referrerID := r.FormValue("referrer_id")

	if referrerID == id {
		http.Error(w, "Cannot refer yourself", http.StatusBadRequest)
		return
	}

	var user, referrer models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := database.DB.First(&referrer, "id = ?", referrerID).Error; err != nil {
		http.Error(w, "Referrer not found", http.StatusNotFound)
		return
	}

	user.ReferrerID = &referrerID

	if err := database.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to update referrer", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "Referrer set successfully",
	})
}
