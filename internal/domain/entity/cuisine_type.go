package entity

import (
	"time"

	"gorm.io/gorm"
)

type CuisineType struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"unique;not null"`
	Description string
	Restaurants []Restaurant   `gorm:"many2many:restaurant_cuisine_types;"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
