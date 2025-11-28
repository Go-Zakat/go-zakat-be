package entity

import "time"

type DonationReceiptItem struct {
	ID          string    `json:"id"`
	ReceiptID   string    `json:"receiptID"`
	FundType    string    `json:"fundType"`    // zakat, infaq, sadaqah
	ZakatType   *string   `json:"zakatType"`   // fitrah, maal (nullable)
	PersonCount *int      `json:"personCount"` // nullable
	Amount      float64   `json:"amount"`
	RiceKG      *float64  `json:"riceKG"` // nullable
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
