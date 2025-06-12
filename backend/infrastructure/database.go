package infrastructure

import (
	"log"
	"path/filepath"

	"github.com/OGARCIIA/examen-backend/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	absPath, err := filepath.Abs(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error getting absolute path for .env: ", err)
	}

	err = godotenv.Load(absPath)
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
