package exception

import (
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/web"
	"net/http"
)

func ErrorHandle(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if badRequestError(writer, request, err) {
		return
	}
	if errorValidate(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}
	if unauthorizedError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponseError := web.ResponseError{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponseError)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponseError := web.ResponseError{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Error:  exception.Error,
		}

		helper.TextFailed("404 NOT FOUND", exception.Error)
		helper.WriteToResponseBody(writer, webResponseError)
		return true
	} else {
		return false
	}
}

func errorValidate(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(ValidationError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponseError := web.ResponseError{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Error:  exception.Error,
		}

		helper.TextFailed("400 BAD REQUEST", exception.Error)
		helper.WriteToResponseBody(writer, webResponseError)
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponseError := web.ResponseError{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Error:  exception.Error,
		}

		helper.TextFailed("400 BAD REQUEST", exception.Error)
		helper.WriteToResponseBody(writer, webResponseError)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponseError := web.ResponseError{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err,
	}

	helper.TextFailed("500 INTERNAL SERVER ERROR", err)
	helper.WriteToResponseBody(writer, webResponseError)

}
