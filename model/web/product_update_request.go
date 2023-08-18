package web

type ProductUpdateRequest struct {
	IdCategory  int     `validate:"required,number,min=1" json:"id_category"`
	SKU         string  `validate:"required,min=1,max=100" json:"sku"`
	Name        string  `validate:"required,min=1,max=100" json:"name"`
	Description string  `json:"description"`
	Price       float64 `validate:"required,number,min=1" json:"price"`
	Stock       int     `validate:"required,number,min=1" json:"stock"`
	Brand       string  `validate:"required,min=1,max=100" json:"brand"`
	Weight      float64 `json:"weight"`
	Dimension   string  `validate:"required,min=1,max=20" json:"dimension"`
	Variant     string  `validate:"required,min=1,max=100" json:"variant"`
	ModifiedBy  string  `validate:"required,min=1,max=100" json:"modified_by"`
}
