package db

import (
	"bjm/db/benjamit"
	"bjm/db/benjamit/models"
	"log"

	"gorm.io/gorm"
)

func Migrate() {
	db, _ := benjamit.Connect()
	if err := createUUIDExtension(db); err != nil {
		log.Fatal("failed to create uuid-ossp extension:", err)
		return
	}

	err := db.AutoMigrate(
		&models.Prefix{},
		&models.User{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("The migration is complete.")
}

func createUUIDExtension(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
}
