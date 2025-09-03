package handler

import (
	"auth-api/internal/database"
	"auth-api/internal/model"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var input model.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// ✅ Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// ✅ Save new user
	user := model.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	// ✅ Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
