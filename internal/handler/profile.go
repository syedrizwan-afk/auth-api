package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key")

func Profile(w http.ResponseWriter, r *http.Request) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Expect header in format: "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	// Parse the token
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Check expiration
	if exp, ok := (*claims)["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}
	}

	// Extract email from claims
	email, ok := (*claims)["email"].(string)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	// Respond with profile data
	profileData := map[string]string{
		"email": email,
		"role":  "user", // You can replace this with actual role from DB if needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profileData)

	fmt.Println("Profile accessed for:", email)

		userID := r.Context().Value("userID").(uint)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message" :  "welcome to your profile",
			"user_id" : userID, 
		})
}
