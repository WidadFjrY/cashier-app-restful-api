package domain

type Products struct {
	Id          int
	SKU         string
	Name        string
	Description string
	Price       float64
	Stock       int
	Brand       string
	Weight      float64
	Dimension   string
	Variant     string
	ModifiedBy  string
	AddedBy     string
	CreatedAt   string
	UpdatedAt   string
}

type ProductCreateOrUpdate struct {
	Id          int
	SKU         string
	Name        string
	Description string
	Price       float64
	Stock       int
	Brand       string
	Weight      float64
	Dimension   string
	Variant     string
	ModifiedBy  string
	AddedBy     string
}

type ProductJoinWithCategory struct {
	Id                 int
	SKU                string
	Name               string
	Description        string
	Price              float64
	Stock              int
	Brand              string
	Weight             float64
	Dimension          string
	Variant            string
	AddedBy            string
	ModifiedBy         string
	CreatedAt          string
	UpdatedAt          string
	IdCategory         int
	NameCategory       string
	CategoryAddedBy    string
	CategoryModifiedBy string
	CategoryCreatedAt  string
	CategoryUpdatedAt  string
}
