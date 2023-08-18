package web

type CategoryUpdateRequest struct {
	ProductId int    `validate:"required,min=1" json:"product_id"`
	Name      string `validate:"required,min=1,max=100" json:"name"`
	AddedBy   string `validate:"required,min=1,max=100" json:"added_by"`
}
