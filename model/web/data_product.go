package web

type CreateDataProduct struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Brand       string  `json:"brand"`
	Weight      float64 `json:"weight"`
	Dimension   string  `json:"dimension"`
	Variant     string  `json:"variant"`
	AddedBy     string  `json:"added_by"`
}

type UpdateDataProduct struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Brand       string  `json:"brand"`
	Weight      float64 `json:"weight"`
	Dimension   string  `json:"dimension"`
	Variant     string  `json:"variant"`
	ModifiedBy  string  `json:"modified_by"`
}

type GetDataProduct struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Brand       string  `json:"brand"`
	Weight      float64 `json:"weight"`
	Dimension   string  `json:"dimension"`
	Variant     string  `json:"variant"`
	ModifiedBy  string  `json:"modified_by"`
	AddedBy     string  `json:"added_by"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
