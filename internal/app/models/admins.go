package models

type Admin struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID int  `json:"user_id" gorm:"not null"`
	User   User `json:"user" gorm:"foreignkey:UserID"`
}
