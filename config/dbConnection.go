package config

import (
	"fmt"
	"os"

	"github.com/MatiasGalli/MS_Product/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var dbURL = os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL environment variable missing")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	autoMigrate(DB)

}

func autoMigrate(connection *gorm.DB) error {

	modelsToMigrate := []interface{}{
		&models.Product{},
		&models.Category{},
	}

	if err := connection.AutoMigrate(modelsToMigrate...); err != nil {
		fmt.Println("Error during AutoMigrate:", err)
		return err
	}

	fmt.Println("AutoMigrate completed successfully")
	return nil
}
