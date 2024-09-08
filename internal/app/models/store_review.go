package models

import (
	"gorm.io/gorm"
	"time"
)

type StoreReview struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	StoreID   uint           `json:"store_id" gorm:"not null"`
	Store     Store          `json:"-" gorm:"foreignKey:StoreID"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	User      User           `json:"-" gorm:"foreignKey:UserID"`
	Rating    uint           `json:"rating" gorm:"not null"`
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
