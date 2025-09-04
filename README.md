# Auth API (Go + MySQL)

A simple authentication API built with **Golang**, **MySQL**, and **JWT**.  
This project implements **user registration, login, and profile retrieval** with JWT-based authentication.

---

## 🚀 Features
- User Registration with hashed passwords (bcrypt)
- User Login with JWT token generation
- Protected Profile endpoint (JWT required)
- MySQL database integration
- Environment variable configuration

---

## 📂 Project Structure
auth-api/
├── cmd/
│ └── main.go # App entry point
├── internal/
│ ├── handler/ # HTTP handlers (register, login, profile)
│ ├── model/ # DB models and types
│ ├── config/ # .env loader
│ └── database/ # DB connection logic
├── go.mod
├── .env # Environment variables (ignored by git)
├── .gitignore
└── README.md
---

## ⚙️ Setup Instructions

### 1. Clone the repository
```sh
git clone https://github.com/your-username/auth-api.git
cd auth-api

# Install dependency
go mod tidy

#Tech Stack

Go (net/http, JWT, bcrypt)
MySQL
Gorilla Mux (router)
dotenv for environment variables
