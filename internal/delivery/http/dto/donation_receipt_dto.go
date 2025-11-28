package dto

import "time"

// Request DTOs
type CreateDonationReceiptItemRequest struct {
	FundType    string   `json:"fund_type" binding:"required,oneof=zakat infaq sadaqah"`
	ZakatType   *string  `json:"zakat_type" binding:"omitempty,oneof=fitrah maal"`
	PersonCount *int     `json:"person_count" binding:"omitempty,min=1"`
	Amount      float64  `json:"amount" binding:"required,gt=0"`
	RiceKG      *float64 `json:"rice_kg" binding:"omitempty,gt=0"`
	Notes       string   `json:"notes"`
}

type CreateDonationReceiptRequest struct {
	MuzakkiID     string                             `json:"muzakki_id" binding:"required"`
	ReceiptNumber string                             `json:"receipt_number" binding:"required"`
	ReceiptDate   string                             `json:"receipt_date" binding:"required"` // YYYY-MM-DD
	PaymentMethod string                             `json:"payment_method" binding:"required"`
	Notes         string                             `json:"notes"`
	Items         []CreateDonationReceiptItemRequest `json:"items" binding:"required,min=1,dive"`
}

type UpdateDonationReceiptRequest struct {
	MuzakkiID     string                             `json:"muzakki_id" binding:"required"`
	ReceiptNumber string                             `json:"receipt_number" binding:"required"`
	ReceiptDate   string                             `json:"receipt_date" binding:"required"`
	PaymentMethod string                             `json:"payment_method" binding:"required"`
	Notes         string                             `json:"notes"`
	Items         []CreateDonationReceiptItemRequest `json:"items" binding:"required,min=1,dive"`
}

// Response DTOs
type DonationReceiptItemResponse struct {
	ID          string   `json:"id"`
	FundType    string   `json:"fund_type"`
	ZakatType   *string  `json:"zakat_type"`
	PersonCount *int     `json:"person_count"`
	Amount      float64  `json:"amount"`
	RiceKG      *float64 `json:"rice_kg"`
	Notes       string   `json:"notes"`
}

type MuzakkiInfo struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type UserInfo struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type DonationReceiptResponse struct {
	ID            string                        `json:"id"`
	ReceiptNumber string                        `json:"receipt_number"`
	ReceiptDate   string                        `json:"receipt_date"`
	Muzakki       MuzakkiInfo                   `json:"muzakki"`
	PaymentMethod string                        `json:"payment_method"`
	TotalAmount   float64                       `json:"total_amount"`
	Notes         string                        `json:"notes"`
	CreatedByUser UserInfo                      `json:"created_by_user"`
	Items         []DonationReceiptItemResponse `json:"items"`
	CreatedAt     time.Time                     `json:"created_at"`
	UpdatedAt     time.Time                     `json:"updated_at"`
}

// List item response (simplified)
type DonationReceiptListItemResponse struct {
	ID              string    `json:"id"`
	ReceiptNumber   string    `json:"receipt_number"`
	ReceiptDate     string    `json:"receipt_date"`
	MuzakkiID       string    `json:"muzakki_id"`
	MuzakkiName     string    `json:"muzakki_name"`
	PaymentMethod   string    `json:"payment_method"`
	TotalAmount     float64   `json:"total_amount"`
	Notes           string    `json:"notes"`
	CreatedByUserID string    `json:"created_by_user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
