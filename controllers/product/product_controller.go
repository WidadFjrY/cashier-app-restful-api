package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
