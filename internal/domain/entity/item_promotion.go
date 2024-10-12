package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemPromotion struct {
	ID             uuid.UUID      `gorm:"type:uuid;" json:"id"`
	MenuItemID     uint64         `gorm:"not null" json:"menu_item_id"`
	Title          string         `json:"title"`           // Title of the promotion
	Description    *string        `json:"description"`     // Detailed description of the promotion
	DiscountAmount float64        `json:"discount_amount"` // Discount amount (e.g., in currency)
	DiscountType   string         `json:"discount_type"`   // Type of discount (e.g., "percentage", "fixed")
	StartDate      time.Time      `json:"start_date"`      // Start date of the promotion
	EndDate        time.Time      `json:"end_date"`        // End date of the promotion
	IsActive       bool           `json:"is_active"`       // Status of the promotion (active/inactive)
	CreatedAt      time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
