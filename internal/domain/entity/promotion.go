package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Promotion struct {
	ID            uuid.UUID      `gorm:"type:uuid;" json:"id"` // Unique identifier for the promotion
	RestaurantID  string         `json:"restaurant_id"`        // Identifier for the associated restaurant
	Title         string         `json:"title"`                // Title of the promotion
	Description   string         `json:"description"`          // Detailed description of the promotion
	PromotionType string         `json:"promotion_type"`       // Type of promotion (e.g., "buy X, get Y", "percentage off")
	RequiredItems *int           `json:"required_items"`       // Number of items that must be purchased
	FreeItems     *int           `json:"free_items"`           // Number of free items received
	Price         float64        `json:"price"`                // Price for the promotion (if applicable)
	StartDate     time.Time      `json:"start_date"`           // Start date of the promotion
	EndDate       time.Time      `json:"end_date"`             // End date of the promotion
	DayOfWeek     string         `json:"day_of_week"`          // Day of the week for the promotion
	IsActive      bool           `json:"is_active"`            // Status of the promotion (active/inactive)
	CreatedAt     time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
