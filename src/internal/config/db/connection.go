package db

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Singleton pattern
var instanceDB *gorm.DB

func GetConnectDB() (*gorm.DB, error) {

	if instanceDB != nil {
		fmt.Println("Using the same connection")
		return instanceDB, nil
	}

	dbUri := os.Getenv("DATABASE_URL")

	instanceDB, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := instanceDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)

	return instanceDB, nil
}
