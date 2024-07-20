package database

import (
	"log"

	"product-management-system/pkg/model"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := TESTConnectDB()
	db.AutoMigrate(
		&model.Product{},
	)
	log.Printf("Migrate: Success")
}
