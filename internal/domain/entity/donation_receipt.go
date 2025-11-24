package entity

import "time"

type DonationReceipt struct {
	ID             string    `json:"id"`
	MuzakkiID      string    `json:"MuzakkiID"`
	ReceiptNumber  string    `json:"receiptNumber"`
	PaymentMethod  string    `json:"paymentMethod"`
	TotalAmount    string    `json:"totalAmount"`
	Notes          string    `json:"notes"`
	CreateByUserID string    `json:"createByUserID"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
