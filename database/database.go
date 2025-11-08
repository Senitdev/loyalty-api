package database

import (
	"fmt"
	"log"
	"loyalty-api/internal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Charger le fichier .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ Pas de fichier .env trouvé, on utilise les variables d'environnement du système")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s  password=%s  dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("adresse ip : ENV  %s %s  %s", os.Getenv("APP_ENV"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
	}
	// Migration des tables
	if err := db.AutoMigrate(&models.User{}, &models.Merchant{}, &models.LoyaltyCard{}, &models.Redemption{}, &models.Reward{}, &models.Transaction{}, &models.Clients{}); err != nil {
		panic("AutoMigrate failed: " + err.Error())
	}
	fmt.Println("Connexion à la base de données réussi")
	DB = db
}
