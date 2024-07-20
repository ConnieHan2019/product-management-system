package database

import (
	"log"

	"gorm.io/gorm"

	"product-management-system/pkg/model"
)

func Migrate(db *gorm.DB) {
	log.Printf("Migrate: Start")

	db.AutoMigrate(
		&model.Product{},
	)
	log.Printf("Migrate: Success")
}
