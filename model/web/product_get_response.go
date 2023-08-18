package web

type ProductGetResponse struct {
	IdProduct   int            `json:"id_product"`
	DataProduct GetDataProduct `json:"data_product"`
	Categories  []Categories   `json:"categories"`
}

type ProductGetResponseWithPages struct {
	Items []ProductGetResponse `json:"items"`
	Pages Pages                `json:"pages"`
}

type Categories struct {
	IdCategory   int             `json:"id_category"`
	DataCategory GetDataCategory `json:"data_category"`
}
