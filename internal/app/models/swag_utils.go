package models

// TokenResponse represents the response with access token and user ID
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
}

// ErrorResponse represents an error message response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DefaultResponse represents a default message response
type DefaultResponse struct {
	Error string `json:"error"`
}

type UserRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	HashPassword string `json:"password"`
}

type UserLogin struct {
	Username     string `json:"username"`
	HashPassword string `json:"password"`
}

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"` // Название категории, обязательное поле
	ParentID     uint   `json:"parent_id,omitempty"`              // Идентификатор родительской категории, необязательное поле
	Description  string `json:"description,omitempty"`            // Описание категории, необязательное поле
}

type AddressRequest struct {
	AddressName string `json:"address_name"`
}

type AccountRequest struct {
	AccountName string `json:"account_number"`
}

type FillAccountRequest struct {
	AccountName string `json:"account_number"`
	Balance     uint   `json:"balance"`
}

type AccountsResponse struct {
	AccountName string `json:"account_name"`
}
