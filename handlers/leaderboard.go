package handlers

import (
	"encoding/json"
	"net/http"

	"user-management-server/models"
	"user-management-server/services"
)

func GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetTopUsersByPoints(10)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var leaderboard []models.LeaderboardUser
	for _, user := range users {
		leaderboard = append(leaderboard, models.LeaderboardUser{
			Name:   user.Name,
			Email:  user.Email,
			Points: user.Points,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leaderboard)
}
