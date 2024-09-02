package models

import (
	"gorm.io/gorm"
	"time"
)

// UserProfile represents a user profile in the system.
type UserProfile struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Age      int    `json:"age"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
}

// Account represents a bank account linked to a user.
type Account struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	User          User           `gorm:"foreignKey:UserID"`
	AccountNumber string         `gorm:"unique;not null" json:"account_number"`
	Balance       float64        `gorm:"default:12100.09" json:"balance"`
	IsDeleted     bool           `gorm:"default:false" json:"is_deleted"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// Address represents a user's address.
type Address struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	AddressName string         `gorm:"size:100;not null" json:"address_name"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID"`
	IsDeleted   bool           `gorm:"default:false" json:"is_deleted"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Category represents a product category.
type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CategoryName string         `gorm:"size:100;not null" json:"category_name"`
	ParentID     *uint          `json:"parent_id,omitempty"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// Product represents a product in the system.
type Product struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	StoreID          uint           `gorm:"not null" json:"store_id"`
	Store            Store          `gorm:"foreignKey:StoreID" json:"store"`
	CategoryID       uint           `gorm:"not null" json:"category_id"`
	Category         Category       `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;"`
	Title            string         `gorm:"size:100;not null" json:"title"`
	Description      string         `gorm:"not null" json:"description"`
	Price            float64        `gorm:"not null" json:"price"`
	Amount           int            `gorm:"not null" json:"amount"`
	DefaultAccountID *uint          `gorm:"default:NULL" json:"default_account_id,omitempty"`
	DefaultAccount   *Account       `gorm:"foreignKey:DefaultAccountID"`
	IsDeleted        bool           `gorm:"default:false" json:"is_deleted"`
	Views            int            `gorm:"default:0" json:"views"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// FeaturedProduct represents a featured product.
type FeaturedProduct struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    uint    `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID"`
	IsDeleted bool    `gorm:"default:false" json:"is_deleted"`
}

// ProductImage represents an image of a product.
type ProductImage struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Image     string  `gorm:"not null" json:"image"`
}

// Review represents a review made by a user on a product.
type Review struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Rating    uint      `gorm:"not null;check:rating >= 1 and rating <= 5" json:"rating"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
}

// Comment represents a comment made by a user on a product.
type Comment struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID"`
	ProductID   uint           `gorm:"not null" json:"product_id"`
	Product     Product        `gorm:"foreignKey:ProductID"`
	ParentID    *uint          `json:"parent_id,omitempty"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// OrderStatus represents the status of an order.
type OrderStatus struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	StatusName  string         `gorm:"size:100;not null" json:"status_name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// OrderDetails represents the details of an order.
type OrderDetails struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `gorm:"foreignKey:ProductID"`
	Price     *float64       `json:"price,omitempty"`
	Quantity  int            `gorm:"default:1" json:"quantity"`
	IsDeleted bool           `gorm:"default:false" json:"is_deleted"`
	OrderDate time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"order_date"`
	AddressID uint           `gorm:"not null" json:"address_id"`
	Address   Address        `gorm:"foreignKey:AddressID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Order represents a user's order.
type Order struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	User         User           `gorm:"foreignKey:UserID"`
	StatusID     uint           `gorm:"not null" json:"status_id"`
	Status       OrderStatus    `gorm:"foreignKey:StatusID"`
	OrderDetails []OrderDetails `gorm:"many2many:order_orderdetails;"` // Изменено на many2many
	IsPaid       *bool          `gorm:"default:false" json:"is_paid"`
	IsInTheCard  bool           `gorm:"default:true" json:"is_in_the_card"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// Payment represents a payment made by a user.
type Payment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	Order     Order          `gorm:"foreignKey:OrderID"`
	Amount    int            `gorm:"not null" json:"amount"`
	Price     float64        `gorm:"not null" json:"price"`
	AccountID uint           `gorm:"not null" json:"account_id"`
	Account   Account        `gorm:"foreignKey:AccountID"`
	PayedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"payed_at"`
	IsDeleted bool           `gorm:"default:false" json:"is_deleted"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
