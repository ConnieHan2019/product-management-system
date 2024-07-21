package request

import (
	"fmt"
	"strings"
)

// Product(dto) is a struct that represents the product entity.
type Product struct {
	UUID         string  `json:"uuid"`
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	ProductCode  string  `json:"productCode"`
	ProductImage string  `json:"productImage"`
	ProductPlace string  `json:"productPlace"`
	Size         string  `json:"size"`
	Weight       string  `json:"weight"`
	ProviderID   string  `json:"providerId"`
	ProviderName string  `json:"-"`
}

// Validate validates the product entity
func (p *Product) Validate() error {
	if p.ProductName == "" {
		return fmt.Errorf("productName is required")
	}
	p.ProductName = strings.TrimSpace(p.ProductName)
	if p.Price == 0 {
		return fmt.Errorf("price is required")
	}
	if p.Description == "" {
		return fmt.Errorf("description is required")
	}
	if p.Category == "" {
		return fmt.Errorf("category is required")
	}
	p.Category = strings.TrimSpace(p.Category)
	if p.ProductCode == "" {
		return fmt.Errorf("productCode is required")
	}
	if p.Stock == 0 {
		return fmt.Errorf("stock is required")
	}
	return nil

}

// ListProductOptions is a struct that represents the options to list products
type ListProductOptions struct {
	ListFilter
	ProductName string  `json:"productName"`
	Category    string  `json:"category"`
	MinPrice    float64 `json:"minPrice"`
	MaxPrice    float64 `json:"maxPrice"`
	// OnlyAvailable is a flag to filter only  products with stock > 0
	OnlyAvailable bool   `json:"onlyAvailable"`
	ProviderName  string `json:"providerName"`
}
