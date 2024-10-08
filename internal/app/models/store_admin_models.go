package models

import (
	"github.com/lib/pq"
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
	User          User           `json:"-" gorm:"foreignKey:UserID"`
	AccountNumber string         `gorm:"unique;not null" json:"account_number"`
	Balance       float64        `json:"balance"`
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
	User        User           `json:"-" gorm:"foreignKey:UserID"`
	IsDeleted   bool           `gorm:"default:false" json:"is_deleted"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Category represents a product category.
type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CategoryName string         `gorm:"unique;size:100;not null" json:"category_name"`
	ParentID     uint           `json:"parent_id"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// Product represents a product in the system.
type Product struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	StoreID          uint           `gorm:"not null" json:"store_id"`
	Store            Store          `json:"-" gorm:"foreignKey:StoreID"`
	CategoryID       uint           `gorm:"not null" json:"category_id"`
	Category         Category       `json:"-" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;"`
	Title            string         `gorm:"size:100;not null" json:"title"`
	Description      string         `gorm:"not null" json:"description"`
	Price            float64        `gorm:"not null" json:"price"`
	Amount           uint           `gorm:"not null" json:"amount"`
	ProductImageList pq.StringArray `gorm:"type:text[]" json:"product_image"`
	Views            int            `gorm:"default:0" json:"views"`
}

// FeaturedProduct represents a featured product.
type FeaturedProduct struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `json:"-" gorm:"foreignKey:ProductID"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `json:"-" gorm:"foreignKey:UserID"`
	IsDeleted bool           `gorm:"default:false" json:"is_deleted"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// ProductImage represents an image of a product.
type ProductImage struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `json:"-" gorm:"foreignKey:ProductID"`
	Image     string  `gorm:"not null" json:"image"`
}

// Review represents a review made by a user on a product.
type Review struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `json:"-" gorm:"foreignKey:UserID"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `json:"-" gorm:"foreignKey:ProductID"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"not null" json:"content"`
	Rating    uint           `gorm:"not null;check:rating >= 1 and rating <= 5" json:"rating"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Comment represents a comment made by a user on a product.
type Comment struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        User           `json:"-" gorm:"foreignKey:UserID"`
	ProductID   uint           `gorm:"not null" json:"product_id"`
	Product     Product        `json:"-" gorm:"foreignKey:ProductID"`
	ParentID    uint           `json:"parent_id,omitempty"`
	CommentText string         `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CommentTree struct {
	Comment  Comment       `json:"comment"`
	Children []CommentTree `json:"children"`
}

// OrderStatus represents the status of an order.
type OrderStatus struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	StatusName  string         `gorm:"unique;size:100;not null" json:"status_name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// OrderDetails represents the details of an order.
type OrderDetails struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `json:"-" gorm:"foreignKey:ProductID"`
	Price     float64        `json:"price,omitempty"`
	Quantity  uint           `gorm:"default:1" json:"quantity"`
	AddressID uint           `gorm:"not null" json:"address_id"`
	Address   Address        `json:"-" gorm:"foreignKey:AddressID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Order represents a user's order.
type Order struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	UserID         uint           `gorm:"not null" json:"user_id"`
	User           User           `json:"-" gorm:"foreignKey:UserID"`
	StatusID       uint           `gorm:"not null" json:"status_id"`
	Status         OrderStatus    `json:"-" gorm:"foreignKey:StatusID"`
	OrderDetailsID uint           `gorm:"not null" json:"order_details_id"`
	OrderDetails   OrderDetails   `json:"order_details" gorm:"foreignKey:OrderDetailsID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type OrderRequestJsonBind struct {
	UserID    uint `json:"user_id"`
	StatusID  uint `json:"status_id"`
	AddressID uint `json:"address_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

// Payment represents a payment made by a user.
type Payment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `json:"-" gorm:"foreignKey:UserID"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	Order     Order          `json:"-" gorm:"foreignKey:OrderID"`
	Amount    uint           `gorm:"not null" json:"amount"`
	Price     float64        `gorm:"not null" json:"price"`
	AccountID uint           `gorm:"not null" json:"account_id"`
	Account   Account        `json:"-" gorm:"foreignKey:AccountID"`
	IsDeleted bool           `gorm:"default:false" json:"is_deleted"`
	PayedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"payed_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Account) TableName() string {
	return "accountapp_account"
}

func (Order) TableName() string {
	return "orderapp_order"
}

func (OrderDetails) TableName() string {
	return "orderapp_orderdetails"
}

func (OrderStatus) TableName() string {
	return "orderapp_orderstatus"
}

func (FeaturedProduct) TableName() string {
	return "featured_productapp_featuredproduct"
}

func (Payment) TableName() string {
	return "payapp_payment"
}

func (Product) TableName() string {
	return "productapp_product"
}

func (Address) TableName() string {
	return "addressapp_address"
}

func (UserProfile) TableName() string {
	return "userapp_userprofile"
}

func (Review) TableName() string {
	return "reviewapp_review"
}

func (Comment) TableName() string {
	return "commentapp_comment"
}

func (Category) TableName() string {
	return "categoryapp_category"
}

func (ProductImage) TableName() string {
	return "productapp_productimage"
}
