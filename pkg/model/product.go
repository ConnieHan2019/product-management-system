package model

// Product defines the model for product
type Product struct {
	UUID        string  `json:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt   string  `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt   string  `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	ProductName string  `json:"productName" gorm:"column:product_name;index;not null"`
	Description string  `json:"description" gorm:"column:description"`
	Category    string  `json:"category" gorm:"column:category;index;not null"`
	Price       float64 `json:"price" gorm:"column:price"`
	Stock       int     `json:"stock" gorm:"column:stock"`

	ProductCode  string `json:"productCode" gorm:"column:product_code"`
	ProductImage string `json:"productImage" gorm:"column:product_image"`
	ProductPlace string `json:"productPlace" gorm:"column:product_place"`
	Size         string `json:"size" gorm:"column:size"`
	Weight       string `json:"weight" gorm:"column:weight"`

	ProviderID   string `json:"providerId" gorm:"column:provider_id"`
	ProviderName string `json:"-" gorm:"-"`
}

// TableName returns the table name
func (p *Product) TableName() string {
	return "products"
}
