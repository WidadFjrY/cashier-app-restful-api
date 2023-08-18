package web

type UserUpdateOrCreateResponse struct {
	Id   int            `json:"id"`
	Data UpdateDataUser `json:"data"`
}
