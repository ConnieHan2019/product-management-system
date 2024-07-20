package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	UUID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`

	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}

func (m *Model) Db() *gorm.DB {
	return config.Db()
}
