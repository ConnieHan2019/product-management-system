package model


// Catagory defines the model for catagory
type Catagory struct {
	UUID         string `json:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt    string `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt    string `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	CatagoryName string `json:"name" gorm:"column:name;index;not null"`
	Description  string `json:"description" gorm:"column:description"`
	Stock        int    `json:"stock" gorm:"column:stock"`
}

func (c *Catagory) TableName() string {
	return "catagories"
}
