package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SavedRestaurant struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uuid.UUID
	RestaurantID uuid.UUID
	Notes        string
	CreatedAt    time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
