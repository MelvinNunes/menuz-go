package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID             uuid.UUID `gorm:"type:uuid;" json:"id"`
	Name           string    `gorm:"not null"`
	Address        string    `gorm:"not null"`
	PriceRangeDown float64   `gorm:"null" json:"price_range_down"`
	PriceRangeUp   float64   `gorm:"null" json:"price_range_up"`
	Latitude       float64   `gorm:"null"`
	Longitude      float64   `gorm:"null"`
	IsApproved     bool      `gorm:"default:false" json:"is_approved"`
	CreatedBy      uuid.UUID `gorm:"type:uuid;" json:"created_by"`
	CreatedByUser  User      `gorm:"foreignKey:CreatedBy"`
	Menus          []Menu
	Reviews        []Review
	CuisineTypes   []CuisineType  `gorm:"many2many:restaurant_cuisine_types;"`
	CreatedAt      time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
