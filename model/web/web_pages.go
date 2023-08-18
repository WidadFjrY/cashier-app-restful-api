package web

type Pages struct {
	Page       int     `json:"page"`
	TotalPages float64 `json:"total_pages"`
	TotalItems int     `json:"total_items"`
}
