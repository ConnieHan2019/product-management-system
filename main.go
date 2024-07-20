package main

import (
	"fmt"

	"product-management-system/pkg/database"
)

func main() {
	fmt.Println("Start Product Management System")
	database.Migrate()
	database.Seeder()
}
