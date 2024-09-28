package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role" gorm:"type:enum('admin','member')"`
	Verify    bool           `json:"verify"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Contacts  *Contacts      `gorm:"foreignKey:UserID"`
	Blog      *[]Blog        `gorm:"foreignKey:Author"`
}
