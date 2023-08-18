package domain

type Categories struct {
	Id         int
	ProductId  int
	Name       string
	AddedBy    string
	ModifiedBy string
	CreatedAt  string
	UpdatedAt  string
}

type CategoriesUpdateOrCreate struct {
	Id        int
	ProductId int
	Name      string
	AddedBy   string
}
