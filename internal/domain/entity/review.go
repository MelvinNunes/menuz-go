package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID           uuid.UUID      `gorm:"type:uuid;" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;" json:"user_id"`
	RestaurantID uuid.UUID      `gorm:"type:uuid;" json:"restaurant_id"`
	MenuItemID   uint64         `gorm:"not null" json:"menu_item_id"`
	Rating       int            `gorm:"check:rating >= 1 AND rating <= 5"`
	Comment      string         `gorm:"type:text;null" json:"comment"`
	CreatedAt    time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
