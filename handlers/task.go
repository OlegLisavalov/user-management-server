package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"user-management-server/services"
)

func CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	var input struct {
		TaskType string `json:"task_type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := services.CompleteTask(userID, input.TaskType); err != nil {
		http.Error(w, "Unable to complete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Task completed"})
}
