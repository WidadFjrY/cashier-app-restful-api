package web

type GetDataCategory struct {
	ProductId  int    `json:"product_id"`
	Name       string `json:"name"`
	AddedBy    string `json:"added_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateOrUpdateDataCategory struct {
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
	AddedBy   string `json:"added_by"`
}
