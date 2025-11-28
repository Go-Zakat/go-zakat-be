package dto

import "time"

// Request DTOs
type CreateDistributionItemRequest struct {
	MustahiqID string  `json:"mustahiq_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required,gt=0"`
	Notes      string  `json:"notes"`
}

type CreateDistributionRequest struct {
	DistributionDate string                          `json:"distribution_date" binding:"required"` // YYYY-MM-DD
	ProgramID        *string                         `json:"program_id"`                           // optional
	SourceFundType   string                          `json:"source_fund_type" binding:"required,oneof=zakat_fitrah zakat_maal infaq sadaqah"`
	Notes            string                          `json:"notes"`
	Items            []CreateDistributionItemRequest `json:"items" binding:"required,min=1,dive"`
}

type UpdateDistributionRequest struct {
	DistributionDate string                          `json:"distribution_date" binding:"required"`
	ProgramID        *string                         `json:"program_id"`
	SourceFundType   string                          `json:"source_fund_type" binding:"required,oneof=zakat_fitrah zakat_maal infaq sadaqah"`
	Notes            string                          `json:"notes"`
	Items            []CreateDistributionItemRequest `json:"items" binding:"required,min=1,dive"`
}

// Response DTOs
type DistributionItemResponse struct {
	ID           string  `json:"id"`
	MustahiqID   string  `json:"mustahiq_id"`
	MustahiqName string  `json:"mustahiq_name"`
	AsnafName    string  `json:"asnaf_name"`
	Address      string  `json:"address"`
	Amount       float64 `json:"amount"`
	Notes        string  `json:"notes"`
}

type ProgramInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DistributionResponse struct {
	ID               string                     `json:"id"`
	DistributionDate string                     `json:"distribution_date"`
	Program          *ProgramInfo               `json:"program,omitempty"`
	SourceFundType   string                     `json:"source_fund_type"`
	TotalAmount      float64                    `json:"total_amount"`
	Notes            string                     `json:"notes"`
	CreatedByUser    UserInfo                   `json:"created_by_user"`
	Items            []DistributionItemResponse `json:"items"`
	CreatedAt        time.Time                  `json:"created_at"`
	UpdatedAt        time.Time                  `json:"updated_at"`
}

// List item response (simplified with beneficiary_count)
type DistributionListItemResponse struct {
	ID               string    `json:"id"`
	DistributionDate string    `json:"distribution_date"`
	ProgramID        *string   `json:"program_id,omitempty"`
	ProgramName      string    `json:"program_name,omitempty"`
	SourceFundType   string    `json:"source_fund_type"`
	TotalAmount      float64   `json:"total_amount"`
	BeneficiaryCount int64     `json:"beneficiary_count"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
