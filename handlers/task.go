package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management-server/services"

	"github.com/gorilla/mux"
)

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	taskName := r.FormValue("task")

	if err := services.CompleteTask(id, taskName); err != nil {
		http.Error(w, "Unable to complete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "Task completed"})
}
