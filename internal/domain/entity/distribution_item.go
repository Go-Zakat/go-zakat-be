package entity

import "time"

type DistributionItem struct {
	ID             string    `json:"id"`
	DistributionID string    `json:"distributionID"`
	MustahiqID     string    `json:"mustahiqID"`
	Mustahiq       *Mustahiq `json:"mustahiq,omitempty"`
	Amount         float64   `json:"amount"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
