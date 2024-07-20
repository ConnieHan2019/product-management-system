package service

import (
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"gorm.io/gorm"

	"product-management-system/pkg/core"
	"product-management-system/pkg/dtos"
	"product-management-system/pkg/model"
)

type ProductService struct {
	DB     *gorm.DB
	Logger logr.Logger
}

func NewProductService(log logr.Logger, db *gorm.DB) *ProductService {
	return &ProductService{
		Logger: log,
		DB:     db,
	}
}

// CreateProduct creates a new product
// search product by name before creating
// if product already exists, return error
// if product does not exist, create it
func (ps *ProductService) CreateProduct(productEntity *dtos.Product) (*model.Product, error) {
	// check if product already exists
	if _, err := ps.GetProductByName(productEntity.ProductName); err == nil {
		ps.Logger.Info("Product already exists", "product name", productEntity.ProductName)
		return nil, fmt.Errorf("product already exists:  product name :%s", productEntity.ProductName)
	}

	product := ps.convertProductToModel(productEntity)
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	res := ps.DB.Create(&product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error creating product", "roduct name", product.ProductName)
		return nil, fmt.Errorf("error creating product: %w, product name:%s", res.Error, product.ProductName)

	}
	return product, nil
}

// GetProductByUUID retrieves a product by its UUID
func (ps *ProductService) GetProductByUUID(uuid string) (*dtos.Product, error) {
	product := &model.Product{}
	res := ps.DB.Where("uuid = ?", uuid).First(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error retrieving product", "uuid", uuid)
		return nil, fmt.Errorf("error retrieving product %w, uuid:%s", res.Error, uuid)
	}
	return ps.convertProductToDTO(product), nil
}

// GetProductByName retrieves a product by its name
func (ps *ProductService) GetProductByName(name string) (*dtos.Product, error) {
	product := &model.Product{}
	//fuzzy search by name
	res := ps.DB.Where("product_name = ?", name).First(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error retrieving product", "name", name)
		return nil, fmt.Errorf("error retrieving product: %w, name:%s", res.Error, name)
	}
	return ps.convertProductToDTO(product), nil
}

// ListProducts retrieves a list of products by ProductListOptions
func (ps *ProductService) ListProducts(options *dtos.ListProductOptions) ([]*dtos.Product, error) {
	products := []*model.Product{}
	query := ps.DB
	if options.ProductName != "" {
		// fuzzy search by name
		query = query.Where("product_name LIKE ?", "%"+options.ProductName+"%")
	}
	if options.Category != "" {
		query = query.Where("category = ?", options.Category)
	}
	if options.MinPrice > 0 {
		query = query.Where("price >= ?", options.MinPrice)
	}
	if options.MaxPrice > 0 {
		query = query.Where("price <= ?", options.MaxPrice)
	}
	if options.OnlyAvailable {
		query = query.Where("stock > 0")
	}
	if options.ProviderName != "" {
		query = query.Where("provider_name = ?", options.ProviderName)
	}

	// pagination
	if options.Page > 0 && options.PageSize > 0 {
		query = query.Offset((options.Page - 1) * options.PageSize).Limit(options.PageSize)
	}
	res := query.Find(&products)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error listing products", "options", options)
		return nil, fmt.Errorf("error listing products: %w, options:%v", res.Error, options)
	}
	productDTOs := []*dtos.Product{}
	for _, product := range products {
		productDTOs = append(productDTOs, ps.convertProductToDTO(product))
	}
	return productDTOs, nil
}

// // UpdateProduct updates a product
// search product by UUID before updating
// if product does not exist, return error
// if product exists, update it
func (ps *ProductService) UpdateProduct(productEntity *dtos.Product) (*model.Product, error) {
	product := ps.convertProductToModel(productEntity)
	// check if product exists
	if _, err := ps.GetProductByUUID(product.UUID); err != nil {
		ps.Logger.Info("Product does not exist", "product UUID", product.UUID)
		return nil, fmt.Errorf("product does not exist: %w, product UUID:%s", err, product.UUID)
	}

	product.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	res := ps.DB.Save(&product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error updating product", "product name", product.ProductName)
		return nil, fmt.Errorf("error updating product: %w, product name:%s", res.Error, product.ProductName)
	}
	return product, nil
}

// DeleteProductByUUID deletes a product by its UUID
// Truely delete the product from database
func (ps *ProductService) DeleteProductByUUID(uuid string) error {
	product := &model.Product{}
	// check if product exists
	if _, err := ps.GetProductByUUID(uuid); err != nil {
		ps.Logger.Info("Product does not exist", "product UUID", uuid)
		return fmt.Errorf("product does not exist: %w, product UUID:%s", err, uuid)
	}
	res := ps.DB.Where("uuid = ?", uuid).Delete(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error deleting product", "product UUID", uuid)
		return fmt.Errorf("error deleting product: %w, product UUID:%s", res.Error, uuid)
	}
	return nil
}

// DeleteProductByName deletes a product by its name
// Truely delete the product from database
func (ps *ProductService) DeleteProductByName(name string) error {
	product := &model.Product{}
	// check if product exists
	if _, err := ps.GetProductByName(name); err != nil {
		ps.Logger.Info("Product does not exist", "product name", name)
		return fmt.Errorf("product does not exist: %w, product name:%s", err, name)
	}
	res := ps.DB.Where("product_name = ?", name).Delete(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error deleting product", "product name", name)
		return fmt.Errorf("error deleting product: %w, product name:%s", res.Error, name)
	}
	return nil
}

// convertProductToDTO converts a product model to a product DTO
func (ps *ProductService) convertProductToDTO(product *model.Product) *dtos.Product {
	if product == nil {
		return nil
	}
	return &dtos.Product{
		UUID:         product.UUID,
		ProductName:  product.ProductName,
		Price:        product.Price,
		Description:  product.Description,
		Category:     product.Category,
		Stock:        product.Stock,
		ProductCode:  product.ProductCode,
		ProductImage: product.ProductImage,
		ProductPlace: product.ProductPlace,
		Size:         product.Size,
		Weight:       product.Weight,
		ProviderID:   product.ProviderID,
		ProviderName: product.ProviderName,
	}
}

func (ps *ProductService) convertProductToModel(product *dtos.Product) *model.Product {
	if product == nil {
		return nil
	}
	if product.UUID == "" {
		product.UUID = core.GenerateUUID()
	}
	return &model.Product{
		UUID:         product.UUID,
		ProductName:  product.ProductName,
		Price:        product.Price,
		Description:  product.Description,
		Category:     product.Category,
		Stock:        product.Stock,
		ProductCode:  product.ProductCode,
		ProductImage: product.ProductImage,
		ProductPlace: product.ProductPlace,
		Size:         product.Size,
		Weight:       product.Weight,
		ProviderID:   product.ProviderID,
		ProviderName: product.ProviderName,
	}
}
