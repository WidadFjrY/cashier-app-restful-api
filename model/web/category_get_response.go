package web

type CategoryResponse struct {
	Id   int             `json:"id"`
	Data GetDataCategory `json:"data"`
}
