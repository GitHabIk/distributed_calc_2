package main

import (
	"log"
	"net/http"

	"distributed-calculator/internal/auth"
	"distributed-calculator/internal/db"
	"distributed-calculator/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	err := db.InitDB("calculator.db")
	if err != nil {
		log.Fatalf("failed to init DB: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/v1/login", handlers.LoginHandler).Methods("POST")

	secured := r.PathPrefix("/api/v1").Subrouter()
	secured.Use(auth.JWTMiddleware)
	secured.HandleFunc("/calculate", handlers.CalculateHandler).Methods("POST")

	log.Println("HTTP server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
