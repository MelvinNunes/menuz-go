package entity

import (
	"time"

	"gorm.io/gorm"
)

type DietaryInfo struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"unique;not null"`
	Description string
	MenuItems   []MenuItem     `gorm:"many2many:menu_item_dietary_infos;"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
