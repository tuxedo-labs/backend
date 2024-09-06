package entity

import (
	"time"

	"github.com/google/uuid"
)

type Posts struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text;" json:"content"`
	Files     string    `gorm:"type:varchar(255)" json:"files"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
