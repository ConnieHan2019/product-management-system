package model

import "product-management-system/pkg/core"

type Catagory struct {
	core.Model
	CatagoryName string `json:"name" gorm:"column:name;index;not null"`
	Description  string `json:"description" gorm:"column:description"`
	Stock        int    `json:"stock" gorm:"column:stock"`
}

func (c *Catagory) TableName() string {
	return "catagories"
}
