package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	RoleID    uint64         `gorm:"not null" json:"role_id"`
	Role      Role           `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
