package model

import "product-management-system/pkg/core"

type User struct {
	core.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
