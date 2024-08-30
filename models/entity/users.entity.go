package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role" gorm:"type:enum('admin','member')"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Contacts []Contacts `gorm:"foreignKey:UserID"`
}
