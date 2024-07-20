package dtos

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

// ListProductOptions is a struct that represents the options to list products
type ListProductOptions struct {
	ListFilter
	ProductName string  `json:"productName"`
	Catagory    string  `json:"category"`
	MinPrice    float64 `json:"minPrice"`
	MaxPrice    float64 `json:"maxPrice"`
	// OnlyAvailable is a flag to filter only  products with stock > 0
	OnlyAvailable bool   `json:"onlyAvailable"`
	ProviderName  string `json:"providerName"`
}
