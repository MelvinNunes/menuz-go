package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuItem struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	MenuID      uuid.UUID
	Name        string `gorm:"not null"`
	Description string
	Price       float64
	Category    []Category    `gorm:"many2many:menu_item_categories;"`    // e.g., "Appetizer", "Main Course", "Dessert"
	DietaryInfo []DietaryInfo `gorm:"many2many:menu_item_dietary_infos;"` // e.g., "Vegetarian", "Vegan", "Gluten-Free"
	Reviews     []Review
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
