package models

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"unique;not null"`
	Description  string         `json:"description"`
	HashPassword string         `json:"-" gorm:"not null"`
	OwnerID      uint           `json:"owner_id" gorm:"not null"`
	Owner        User           `json:"owner" gorm:"foreignKey:OwnerID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
