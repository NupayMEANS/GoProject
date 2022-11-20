package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}
	DB_userName := os.Getenv("DB_USER")
	DB_password := os.Getenv("DB_Password")
	DB_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_userName, DB_password, DB_name)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection failed")
	}
	// DB.AutoMigrate(&UserDetails{})
	fmt.Println("Databse connceted successfully")

}
