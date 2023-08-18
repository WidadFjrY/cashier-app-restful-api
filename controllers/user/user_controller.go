package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Add(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Verification(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	NewPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
