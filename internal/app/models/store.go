package models

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"unique;not null"`
	Description string         `json:"description"`
	OwnerID     uint           `json:"owner_id" gorm:"not null"`
	Owner       User           `json:"-" gorm:"foreignKey:OwnerID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
