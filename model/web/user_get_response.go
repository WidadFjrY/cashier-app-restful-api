package web

type UserResponses struct {
	Id   int         `json:"id"`
	Data GetDataUser `json:"data"`
}

type UserResponseWithPages struct {
	Items []UserResponses `json:"items"`
	Pages `json:"pages"`
}
