package service

import (
	"fmt"

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
		return nil, fmt.Errorf("Product already exists: %v", "product name", productEntity.ProductName)
	}

	product := ps.convertProductToModel(productEntity)
	res := ps.DB.Create(&product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error creating product", "roduct name", product.Name)
		return nil, fmt.Errorf("Error creating product: %v", res.Error, "product name", product.Name)

	}
	return product, nil
}

// GetProductByUUID retrieves a product by its UUID
func (ps *ProductService) GetProductByUUID(uuid string) (*dtos.Product, error) {
	product := &model.Product{}
	res := ps.DB.Where("uuid = ?", uuid).First(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error retrieving product", "uuid", uuid)
		return nil, fmt.Errorf("Error retrieving product: %v", res.Error, "uuid", uuid)
	}
	return ps.convertProductToDTO(product), nil
}

// GetProductByName retrieves a product by its name
func (ps *ProductService) GetProductByName(name string) (*dtos.Product, error) {
	product := &model.Product{}
	//fuzzy search by name
	res := ps.DB.Where("name LIKE ?", "%"+name+"%").First(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error retrieving product", "name", name)
		return nil, fmt.Errorf("Error retrieving product: %v", res.Error, "name", name)
	}
	return ps.convertProductToDTO(product), nil
}

// ListProducts retrieves a list of products by ProductListOptions
func (ps *ProductService) ListProducts(options *dtos.ListProductOptions) ([]*dtos.Product, error) {
	products := []*model.Product{}
	query := ps.DB
	if options.ProductName != "" {
		// fuzzy search by name
		query = query.Where("name LIKE ?", "%"+options.ProductName+"%")
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
	res := query.Find(&products)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error listing products", "options", options)
		return nil, fmt.Errorf("Error listing products: %v", res.Error, "options", options)
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
		return nil, fmt.Errorf("Product does not exist: %v", "product UUID", product.UUID)
	}

	res := ps.DB.Save(&product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error updating product", "product name", product.Name)
		return nil, fmt.Errorf("Error updating product: %v", res.Error, "product name", product.Name)
	}
	return product, nil
}

// DeleteProductByUUID deletes a product by its UUID
func (ps *ProductService) DeleteProductByUUID(uuid string) error {
	product := &model.Product{}
	// check if product exists
	if _, err := ps.GetProductByUUID(uuid); err != nil {
		ps.Logger.Info("Product does not exist", "product UUID", uuid)
		return fmt.Errorf("Product does not exist: %v", "product UUID", uuid)
	}
	res := ps.DB.Where("uuid = ?", uuid).Delete(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error deleting product", "product UUID", uuid)
		return fmt.Errorf("Error deleting product: %v", res.Error, "product UUID", uuid)
	}
	return nil
}

// DeleteProductByName deletes a product by its name
func (ps *ProductService) DeleteProductByName(name string) error {
	product := &model.Product{}
	// check if product exists
	if _, err := ps.GetProductByName(name); err != nil {
		ps.Logger.Info("Product does not exist", "product name", name)
		return fmt.Errorf("Product does not exist: %v", "product name", name)
	}
	res := ps.DB.Where("name = ?", name).Delete(product)
	if res.Error != nil {
		ps.Logger.Error(res.Error, "Error deleting product", "product name", name)
		return fmt.Errorf("Error deleting product: %v", res.Error, "product name", name)
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
		ProductName:  product.Name,
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
