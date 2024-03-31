package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID      `gorm:"type:uuid;" json:"id"`
	Code            string         `gorm:"type:varchar(255);not null;unique" json:"code"`
	Email           string         `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password        string         `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumberCode string         `gorm:"type:varchar(255);not null" json:"phone_number_code"`
	PhoneNumber     string         `gorm:"type:varchar(255);unique" json:"phone_number"`
	RawPhoneNumber  string         `gorm:"type:varchar(255);unique" json:"raw_phone_number"`
	AppLanguage     string         `gorm:"type:varchar(255);not null;default:pt" json:"app_language"`
	Active          bool           `gorm:"default:true" json:"active"`
	CreatedAt       time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
