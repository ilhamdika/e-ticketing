package main

import (
	"e-ticketing/config"
	"e-ticketing/models"
	"e-ticketing/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Card{},
		&models.Terminal{},
		&models.Tarif{},
		&models.Transaction{},
		&models.SyncLog{},
	)

	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", "admin").Count(&count)
	if count == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

		admin := models.User{
			Username:    "admin",
			Password:    string(hashed),
			Name:        "Administrator",
			PhoneNumber: "081234567890",
			Email:       "admin@example.com",
		}

		if err := config.DB.Create(&admin).Error; err != nil {
			log.Println("Failed to seed admin user:", err)
		} else {
			log.Println("Admin user created: admin / admin123")
		}
	}

	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
