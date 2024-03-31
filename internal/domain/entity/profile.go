package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID"`
	FirstName   string         `gorm:"type:varchar(255)" json:"first_name"`
	LastName    string         `gorm:"type:varchar(255)" json:"last_name"`
	Gender      string         `gorm:"type:varchar(255)" json:"gender"`
	Avatar      *string        `gorm:"type:varchar(255);default:null" json:"avatar"`
	DateOfBirth time.Time      `gorm:"type:date" json:"date_of_birth"`
	CreatedAt   time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *Profile) GetFullName() string {
	return fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}
