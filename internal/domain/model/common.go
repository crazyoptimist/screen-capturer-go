package model

import (
	"time"
)

// Same as gorm.Model, but including column and json tags
type Common struct {
	ID        int        `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}
