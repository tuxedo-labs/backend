package entity

import (
	"github.com/google/uuid"
	"time"
)

type Blog struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Thumbnail   string    `gorm:"type:varchar(255)" json:"thumbnail"`
	Author      uint      `json:"author" gorm:"index"`
	User        Users     `gorm:"foreignKey:Author"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
