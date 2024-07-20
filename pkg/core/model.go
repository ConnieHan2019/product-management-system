package core

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	UUID      string    `json:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

// GenerateUUID ...
func GenerateUUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
