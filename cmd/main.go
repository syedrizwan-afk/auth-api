package main

import (
	"fmt"
	"log"
	"net/http"

	//"auth-api/internal/auth"
	"auth-api/internal/config"
	"auth-api/internal/database"
	"auth-api/internal/handler"
	"auth-api/internal/middleware"

	"github.com/gorilla/mux"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from /register")
}

func main() {
	config.LoadEnv()
	database.Connect()

	r := mux.NewRouter()
	// r.HandleFunc("/",handler.Main)
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/profile", handler.Login).Methods("POST")

	r.Handle("/profile", middleware.JWTAuth(http.HandlerFunc(handler.Profile))).Methods("GET")
	//http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Hit /register")
	//	handler.Register(w, r)
	//})

	log.Println("Server running on port :5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
