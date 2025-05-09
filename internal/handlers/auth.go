package handlers

import (
	"encoding/json"
	"net/http"

	"distributed-calculator/internal/auth"
	"distributed-calculator/internal/db"
)

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	hash, err := auth.HashPassword(creds.Password)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	err = db.CreateUser(creds.Login, hash)
	if err != nil {
		http.Error(w, "User exists or db error", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID, hash, err := db.GetUserByLogin(creds.Login)
	if err != nil || !auth.CheckPasswordHash(creds.Password, hash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
