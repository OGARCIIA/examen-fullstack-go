package infrastructure

import (
	"log"

	"github.com/OGARCIIA/examen-backend/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	pathsToTry := []string{
		".env",
		"../.env",
		"../../.env",
	}

	var err error
	for _, p := range pathsToTry {
		err = godotenv.Load(p)
		if err == nil {
			log.Printf("Loaded .env from %s", p)
			break
		}
	}

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	dsn := config.GetDBConnectionString()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to database: ", err)
	}

	DB = database
}
