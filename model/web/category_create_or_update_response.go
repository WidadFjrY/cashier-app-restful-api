package web

type CategoryCreateOrUpdateResponse struct {
	Id   int                        `json:"id"`
	Data CreateOrUpdateDataCategory `json:"data"`
}
