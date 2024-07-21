package database

import (
	"fmt"
	"log"
	"math/rand"
	"product-management-system/pkg/model"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	log.Printf("Seeder: Start")

	category := [5]string{"laptop", "mobile phone", "watch", "headphone", "camera"}
	providerName := [4]string{"apple", "samsung", "xiaomi", "huawei"}
	providerID := [4]string{"1", "2", "3", "4"}

	var productList [50]model.Product

	// 创建 10 条 laptop 类型数据
	for i := 0; i < 10; i++ {
		var price float64
		if i < 3 {
			price = rand.Float64()*1000 + 1000 // 1000 to 2000
		} else if i < 6 {
			price = rand.Float64()*1000 + 2000 // 2000 to 3000
		} else {
			price = rand.Float64()*1000 + 3000 // 3000 to 4000
		}
		productList[i] = model.Product{
			UUID:         GenerateUUID(),
			ProductName:  fmt.Sprintf("Laptop_%d", i+1),
			Price:        price,
			Stock:        rand.Intn(10) + 1,
			Category:     category[0],
			ProviderName: providerName[rand.Intn(4)],
			ProviderID:   providerID[rand.Intn(4)],
			ProductCode:  fmt.Sprintf("LP%d", i+1),
			Description:  fmt.Sprintf("Laptop %d description", i+1),
			Size:         fmt.Sprintf("%dx%dx%d", rand.Intn(30)+10, rand.Intn(20)+5, rand.Intn(10)+3),
			Weight:       fmt.Sprintf("%.2f kg", rand.Float64()*2+1),
		}
		db.Create(&productList[i])
	}

	// 创建 10 条 mobile phone 类型数据
	for i := 10; i < 20; i++ {
		var price float64
		if i < 13 {
			price = rand.Float64()*1000 + 1000 // 1000 to 2000
		} else if i < 16 {
			price = rand.Float64()*1000 + 2000 // 2000 to 3000
		} else {
			price = rand.Float64()*1000 + 3000 // 3000 to 4000
		}
		productList[i] = model.Product{
			UUID:         GenerateUUID(),
			ProductName:  fmt.Sprintf("Mobile_Phone %d", i-9),
			Price:        price,
			Stock:        rand.Intn(10) + 1,
			Category:     category[1],
			ProviderName: providerName[rand.Intn(4)],
			ProviderID:   providerID[rand.Intn(4)],
			ProductCode:  fmt.Sprintf("MP%d", i-9),
			Description:  fmt.Sprintf("Mobile Phone %d description", i-9),
			Size:         fmt.Sprintf("%dx%dx%d mm", rand.Intn(20)+50, rand.Intn(10)+60, rand.Intn(5)+10),
			Weight:       fmt.Sprintf("%.2f g", rand.Float64()*200+100),
		}
		db.Create(&productList[i])
	}

	// 创建 10 条 watch 类型数据
	for i := 20; i < 30; i++ {
		var price float64
		if i < 23 {
			price = rand.Float64()*1000 + 1000 // 1000 to 2000
		} else if i < 26 {
			price = rand.Float64()*1000 + 2000 // 2000 to 3000
		} else {
			price = rand.Float64()*1000 + 3000 // 3000 to 4000
		}
		productList[i] = model.Product{
			UUID:         GenerateUUID(),
			ProductName:  fmt.Sprintf("Watch_%d", i-19),
			Price:        price,
			Stock:        rand.Intn(10) + 1,
			Category:     category[2],
			ProviderName: providerName[rand.Intn(4)],
			ProviderID:   providerID[rand.Intn(4)],
			ProductCode:  fmt.Sprintf("W%d", i-19),
			Description:  fmt.Sprintf("Watch %d description", i-19),
			Size:         fmt.Sprintf("%dx%d mm", rand.Intn(10)+30, rand.Intn(10)+20),
			Weight:       fmt.Sprintf("%.2f g", rand.Float64()*50+20),
		}
		db.Create(&productList[i])
	}

	// 创建 10 条 headphone 类型数据
	for i := 30; i < 40; i++ {
		var price float64
		if i < 33 {
			price = rand.Float64()*1000 + 1000 // 1000 to 2000
		} else if i < 36 {
			price = rand.Float64()*1000 + 2000 // 2000 to 3000
		} else {
			price = rand.Float64()*1000 + 3000 // 3000 to 4000
		}
		productList[i] = model.Product{
			UUID:         GenerateUUID(),
			ProductName:  fmt.Sprintf("Headphone_%d", i-29),
			Price:        price,
			Stock:        rand.Intn(10) + 1,
			Category:     category[3],
			ProviderName: providerName[rand.Intn(4)],
			ProviderID:   providerID[rand.Intn(4)],
			ProductCode:  fmt.Sprintf("H%d", i-29),
			Description:  fmt.Sprintf("Headphone %d description", i-29),
			Size:         fmt.Sprintf("%dx%dx%d mm", rand.Intn(10)+100, rand.Intn(10)+50, rand.Intn(5)+20),
			Weight:       fmt.Sprintf("%.2f g", rand.Float64()*100+50),
		}
		db.Create(&productList[i])
	}

	// 创建 10 条 camera 类型数据
	for i := 40; i < 50; i++ {
		var price float64
		if i < 43 {
			price = rand.Float64()*1000 + 1000 // 1000 to 2000
		} else if i < 46 {
			price = rand.Float64()*1000 + 2000 // 2000 to 3000
		} else {
			price = rand.Float64()*1000 + 3000 // 3000 to 4000
		}
		productList[i] = model.Product{
			UUID:         GenerateUUID(),
			ProductName:  fmt.Sprintf("Camera %d", i-39),
			Price:        price,
			Stock:        rand.Intn(10) + 1,
			Category:     category[4],
			ProviderName: providerName[rand.Intn(4)],
			ProviderID:   providerID[rand.Intn(4)],
			ProductCode:  fmt.Sprintf("C%d", i-39),
			Description:  fmt.Sprintf("Camera %d description", i-39),
			Size:         fmt.Sprintf("%dx%dx%d mm", rand.Intn(20)+100, rand.Intn(20)+50, rand.Intn(10)+30),
			Weight:       fmt.Sprintf("%.2f g", rand.Float64()*500+200),
		}
		db.Create(&productList[i])
	}

	log.Printf("Seeder: Success")
}
