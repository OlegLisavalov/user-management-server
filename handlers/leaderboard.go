package handlers

import (
	"encoding/json"
	"net/http"
	"user-management-server/services"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetTopUsersByPoints(10) // Получение топ-10 пользователей по очкам
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
