package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"product-management-system/pkg/log"
	"product-management-system/pkg/model"
	"product-management-system/pkg/request"
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
		productEntity := &request.Product{
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
		searchRes, err := ps.GetProductByUUID(product.UUID)
		assert.Nil(t, err)
		//assert.Equal(t, searchRes1.ProductName, productEntity.ProductName)
		fmt.Printf("productEntity.ProductName: %s, search Res Product Name: %s\n", productEntity.ProductName, searchRes.ProductName)
	}
	ps.DB.Migrator().DropTable(&model.Product{})

}

// TestCreateAndListProduct 测试创建产品并列出产品
func TestCreateAndListProduct(t *testing.T) {
	ps := setup(t)

	// 创建10个laptop类别的产品
	for i := 0; i < 10; i++ {
		productEntity := &request.Product{
			ProductName: "Laptop " + string(rune(65+i)), // 例如：Laptop A, Laptop B, ...
			Price:       float64(1000 + i*100),          // 价格递增
			Description: "A high performance laptop",
			Category:    "laptop",
			Stock:       100 + i*10, // 库存递增
			// 其他字段根据需要设置
		}

		// 调用CreateProduct函数
		_, err := ps.CreateProduct(productEntity)
		if err != nil {
			t.Errorf("failed to create product: %v", err)
			continue
		}
		assert.Nil(t, err)

	}
	// create 10 camera products
	for i := 0; i < 10; i++ {
		productEntity := &request.Product{
			ProductName: "Camera " + string(rune(65+i)),
			Price:       float64(1000 + i*100),
			Description: "A high performance camera",
			Category:    "camera",
			Stock:       100 + i*10,
		}

		// 调用CreateProduct函数
		_, err := ps.CreateProduct(productEntity)
		if err != nil {
			t.Errorf("failed to create product: %v", err)
			continue
		}
		assert.Nil(t, err)

	}
	// List by category
	allProducts, err2 := ps.ListProducts(&request.ListProductOptions{
		Category: "laptop",
	})
	assert.Nil(t, err2)
	assert.Equal(t, 10, len(allProducts))
	for _, product := range allProducts {
		fmt.Println("product name: ", product.ProductName)
	}
	ps.DB.Migrator().DropTable(&model.Product{})
}

// TestCreateExistingProduct 测试创建已存在的产品
func TestCreateExistingProduct(t *testing.T) {
	ps := setup(t)

	// 创建一个产品
	productEntity := &request.Product{
		ProductName: "Laptop A",
		Price:       1000,
		Description: "A high performance laptop",
		Category:    "laptop",
		Stock:       100,
	}
	_, err := ps.CreateProduct(productEntity)
	assert.Nil(t, err)

	// 再次创建相同的产品
	_, err = ps.CreateProduct(productEntity)
	assert.NotNil(t, err)
	ps.DB.Migrator().DropTable(&model.Product{})
}

// TestDeleteProductByUUID 测试删除产品
func TestDeleteProductByUUID(t *testing.T) {
	ps := setup(t)

	// 创建一个产品
	productEntity := &request.Product{
		ProductName: "Laptop A",
		Price:       1000,
		Description: "A high performance laptop",
		Category:    "laptop",
		Stock:       100,
	}
	product, err := ps.CreateProduct(productEntity)
	assert.Nil(t, err)

	// 删除产品
	err = ps.DeleteProductByUUID(product.UUID)
	assert.Nil(t, err)

	// 查询产品
	_, err = ps.GetProductByUUID(product.UUID)
	assert.NotNil(t, err)
	ps.DB.Migrator().DropTable(&model.Product{})
}
