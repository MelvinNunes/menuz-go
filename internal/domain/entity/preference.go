package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Preference struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uuid.UUID
	CuisineType    []CuisineType  `gorm:"many2many:preference_cuisine_types;"` // e.g., "Italian", "Mexican", "Japanese", "Chinese"
	DietaryInfo    []DietaryInfo  `gorm:"many2many:preference_dietary_infos;"` // e.g., "Vegan", "Gluten-free", "Vegetarian"
	PriceRangeDown float64        `gorm:"null" json:"price_range_down"`
	PriceRangeUp   float64        `gorm:"null" json:"price_range_up"`
	CreatedAt      time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
