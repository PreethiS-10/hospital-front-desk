package config

import (
	"fmt"
	"log"
	"os"
	"github.com/PreethiS-10/fiber-crud/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	dir, _ := os.Getwd()
	fmt.Println("Current working directory:", dir)
	var err error
	err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	// Read from environment variables (you can also hardcode these for testing)
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	err = DB.AutoMigrate(
		&models.Patient{},
		&models.Doctor{},
		&models.Department{},
		&models.Schedule{},
		&models.Appointment{},
	)

	if err != nil {
		log.Fatal(" Failed to migrate models:", err)
	}

	log.Println(" Database migrated successfully")
	log.Println("Connected to PostgreSQL database")
}
