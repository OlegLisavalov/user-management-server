package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user-management-server/database"
	"user-management-server/handlers"
)

func main() {
	database.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/users/{id}/referrer", handlers.CheckReferrer).Methods("POST")
	r.HandleFunc("/users/register", handlers.RegisterUser).Methods("POST")

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
