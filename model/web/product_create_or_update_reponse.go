package web

type ProductCreateResponse struct {
	Id   int               `json:"id"`
	Data CreateDataProduct `json:"data"`
}

type ProductUpdateResponse struct {
	Id   int               `json:"id"`
	Data UpdateDataProduct `json:"data"`
}
