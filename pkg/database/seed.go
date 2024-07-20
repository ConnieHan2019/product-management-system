package database

import (
	"log"

	"product-management-system/pkg/model"
)

func Seeder() {
	log.Printf("Seeder: Start")
	db := ConnectDB()
	db.Create(&model.Product{
		ProductName: "Laptop Lenovo",
		Category:    "Laptop",
		Price:       1000,
	})
	db.Create(&model.Product{
		ProductName: "Laptop HP",
		Category:    "Laptop",
		Price:       2000,
	})
	db.Create(&model.Product{
		ProductName: "Laptop Dell",
		Category:    "Laptop",
		Price:       3000,
	})
	log.Printf("Seeder: Success")

}
