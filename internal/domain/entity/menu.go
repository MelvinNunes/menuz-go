package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID           uuid.UUID `gorm:"type:uuid;" json:"id"`
	RestaurantID uuid.UUID
	Name         string `gorm:"not null"` // e.g., "Lunch", "Dinner"
	Items        []MenuItem
	CreatedAt    time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
