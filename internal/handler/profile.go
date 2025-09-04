package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	email := r.Context().Value("email").(string)

	profileData := map[string]interface{}{
		"message": "welcome to your profile",
		"user_id": userID,
		"email":   email,
		"role":    "user", // later from DB if needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profileData)

	log.Println("Profile accessed for:", email)
}
