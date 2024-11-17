package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user-management-server/database"
	"user-management-server/handlers"
	"user-management-server/middleware"
)

func main() {
	database.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/users/{id}/status", handlers.GetUserStatusHandler).Methods("GET")
	protected.HandleFunc("/users/leaderboard", handlers.GetLeaderboardHandler).Methods("GET")
	protected.HandleFunc("/users/{id}/task/complete", handlers.CompleteTaskHandler).Methods("POST")

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
