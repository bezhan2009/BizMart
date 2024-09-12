package models

import (
	"gorm.io/gorm"
	"time"
)

// User godoc
// @Description Represents a user in the system.
// @Param id path int true "User ID"
// @Param username query string true "User's username"
// @Param email query string true "User's email"
// @Param password query string true "User's hashed password"
type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FirstName    string         `json:"first_name" gorm:"not null"`
	LastName     string         `json:"last_name" gorm:"not null"`
	Username     string         `json:"username" gorm:"unique;not null"`
	Email        string         `json:"email" gorm:"unique;not null"`
	HashPassword string         `json:"password" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
