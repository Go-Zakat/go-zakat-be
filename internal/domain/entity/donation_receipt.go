package entity

import "time"

type DonationReceipt struct {
	ID              string                 `json:"id"`
	MuzakkiID       string                 `json:"muzakkiID"`
	Muzakki         *Muzakki               `json:"muzakki,omitempty"`
	ReceiptNumber   string                 `json:"receiptNumber"`
	ReceiptDate     string                 `json:"receiptDate"` // YYYY-MM-DD
	PaymentMethod   string                 `json:"paymentMethod"`
	TotalAmount     float64                `json:"totalAmount"`
	Notes           string                 `json:"notes"`
	CreatedByUserID string                 `json:"createdByUserID"`
	CreatedByUser   *User                  `json:"createdByUser,omitempty"`
	Items           []*DonationReceiptItem `json:"items,omitempty"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}
