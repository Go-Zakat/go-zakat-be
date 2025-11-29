package dto

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type MetaResponse struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}
