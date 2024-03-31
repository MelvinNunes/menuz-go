package entity

import "time"

type Role struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);unique" json:"name"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
	UpdatedAt time.Time `gorm:"null" json:"updated_at"`
}
