package model

// User defines the model for user
type User struct {
	UUID      string `json:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt string `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt string `json:"updatedAt,omitempty" gorm:"column:updated_at"`

	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (u *User) TableName() string {
	return "users"
}
