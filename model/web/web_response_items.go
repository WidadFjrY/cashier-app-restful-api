package web

type ResponseItems struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Items  interface{} `json:"items"`
}
