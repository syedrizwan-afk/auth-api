package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"auth-api/internal/config"
	"auth-api/internal/model"
)

var DB *gorm.DB

func Connect() {
	user := config.GetEnv("DB_USER")
	pass := config.GetEnv("DB_PASSWORD")
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	name := config.GetEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the User model
	db.AutoMigrate(&model.User{})

	DB = db
	log.Println("Database connection established")
}
