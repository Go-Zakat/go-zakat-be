package dto

import "time"

type CreateProgramRequest struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type UpdateProgramRequest struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type ProgramResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
