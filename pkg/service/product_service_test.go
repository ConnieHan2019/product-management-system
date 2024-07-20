package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"product-management-system/pkg/dtos"
	"product-management-system/pkg/log"
	"product-management-system/pkg/model"
)

// 测试前的准备，初始化数据库和ProductService
func setup(t *testing.T) *ProductService {
	// 这里使用SQLite内存数据库进行测试，实际项目中可以根据需要使用其他数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// 迁移模型到数据库
	db.AutoMigrate(&model.Product{})

	// 创建ProductService实例
	logger := log.NewLogger(log.LogOption{
		Level: zap.DebugLevel,
	})

	productService := NewProductService(logger, db)

	return productService
}

// 测试用例
func TestCreateAndGetProduct(t *testing.T) {
	ps := setup(t)

	// 创建10个laptop类别的产品
	for i := 0; i < 10; i++ {
		productEntity := &dtos.Product{
			ProductName: "Laptop " + string(rune(65+i)), // 例如：Laptop A, Laptop B, ...
			Price:       float64(1000 + i*100),          // 价格递增
			Description: "A high performance laptop",
			Category:    "laptop",
			Stock:       100 + i*10, // 库存递增
			// 其他字段根据需要设置
		}

		// 调用CreateProduct函数
		product, err := ps.CreateProduct(productEntity)
		if err != nil {
			t.Errorf("failed to create product: %v", err)
			continue
		}

		// 断言检查
		assert.Equal(t, productEntity.ProductName, product.ProductName)
		assert.Equal(t, productEntity.Category, product.Category)
		searchRes1, err := ps.GetProductByUUID(product.UUID)
		assert.Nil(t, err)
		assert.Equal(t, searchRes1.ProductName, productEntity.ProductName)
		fmt.Printf("productEntity.ProductName: %s, search Res Product Name: %s\n", productEntity.ProductName, searchRes1.ProductName)

		searchRes2, err := ps.GetProductByName(productEntity.ProductName)
		assert.Nil(t, err)
		assert.Equal(t, searchRes2.ProductName, productEntity.ProductName)

	}

}
