package entity

import "time"

type Distribution struct {
	ID               string              `json:"id"`
	DistributionDate string              `json:"distributionDate"` // YYYY-MM-DD
	ProgramID        *string             `json:"programID"`        // nullable
	Program          *Program            `json:"program,omitempty"`
	SourceFundType   string              `json:"sourceFundType"` // zakat_fitrah, zakat_maal, infaq, sadaqah
	TotalAmount      float64             `json:"totalAmount"`
	Notes            string              `json:"notes"`
	CreatedByUserID  string              `json:"createdByUserID"`
	CreatedByUser    *User               `json:"createdByUser,omitempty"`
	Items            []*DistributionItem `json:"items,omitempty"`
	CreatedAt        time.Time           `json:"createdAt"`
	UpdatedAt        time.Time           `json:"updatedAt"`
}
