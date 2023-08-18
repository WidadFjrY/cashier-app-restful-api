package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/web"
	services "github.com/widadfjry/cashier-app/services/user"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(service services.UserService) UserController {
	return UserControllerImpl{UserService: service}
}

func (controller UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	helper.TextSuccess("Request", "POST /api/v1/users/register")
	userResponse := controller.UserService.Create(request.Context(), userRequest)

	writer.WriteHeader(http.StatusCreated)
	webResponse := web.Response{
		Code:   201,
		Status: "CREATED",
		Data:   userResponse,
	}
	helper.TextSuccess("201 CREATED", "POST /api/v1/users/register")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) Add(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	isAdmin := request.Header.Get("isAdmin")
	username := request.Header.Get("username")

	isAdminBool, err := strconv.ParseBool(isAdmin)
	helper.PanicIfError(err)

	helper.TextSuccess("Request", "POST /api/v1/users/add")
	userResponse := controller.UserService.AddUser(request.Context(), userRequest, isAdminBool, username)

	writer.WriteHeader(http.StatusCreated)
	webResponse := web.Response{
		Code:   201,
		Status: "CREATED",
		Data:   userResponse,
	}
	helper.TextSuccess("201 CREATED", "POST /api/v1/users/add")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userRequest)
	helper.TextSuccess("Request", "POST /api/v1/users/login")

	userResponse := controller.UserService.Login(request.Context(), userRequest)
	userLoginResponse := web.UserLoginResponse{Token: userResponse}

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userLoginResponse,
	}

	helper.TextSuccess("200 OK", "POST /api/v1/users/login")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	userIdInt := helper.StringToInt(userId)

	helper.TextSuccess("Request", "GET /api/v1/users/id/"+userId)
	userResponse := controller.UserService.FindById(request.Context(), userIdInt)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.TextSuccess("200 OK", "GET /api/v1/users/id/"+userId)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	page := params.ByName("page")
	pageInt := helper.StringToInt(page)

	helper.TextSuccess("Request", "GET /api/v1/users/pages/"+page)
	userResponses, pages := controller.UserService.FindAll(request.Context(), pageInt)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data: web.UserResponseWithPages{
			Items: userResponses,
			Pages: pages,
		},
	}

	helper.TextSuccess("200 OK", "GET /api/v1/users/pages/"+page)
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller UserControllerImpl) UpdateById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	username := request.Header.Get("username")
	isAdmin := request.Header.Get("isAdmin")

	isAdminBool, err := strconv.ParseBool(isAdmin)
	helper.PanicIfError(err)

	userIdInt := helper.StringToInt(userId)
	userRequest := web.UserUpdateRequest{}

	helper.TextSuccess("Request", "PUT /api/v1/users/id/"+userId)
	helper.ReadFromRequestBody(request, &userRequest)
	userResponse := controller.UserService.UpdateById(request.Context(), userRequest, userIdInt, username, isAdminBool)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.TextSuccess("200 OK", "PUT /api/v1/users/id/"+userId)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) UpdateByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := request.Header.Get("username")
	userRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	helper.TextSuccess("Request", "PUT /api/v1/users/current/update-data")
	userResponse := controller.UserService.UpdateByUsername(request.Context(), userRequest, username)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.TextSuccess("Request", "PUT /api/v1/users/current/update-data")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) DeleteById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	isAdmin := request.Header.Get("isAdmin")

	isAdminBool, err := strconv.ParseBool(isAdmin)
	helper.PanicIfError(err)

	userIdInt := helper.StringToInt(userId)

	helper.TextSuccess("Request", "DELETE /api/v1/users/id/"+userId)
	controller.UserService.DeleteById(request.Context(), userIdInt, isAdminBool)

	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "user deleted",
	}

	helper.TextSuccess("200 OK", "DELETE /api/v1/users/"+userId)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := request.Header.Get("username")
	userRequest := web.UserChangePasswordRequest{}

	helper.TextSuccess("Request", "PUT /api/v1/users/current/change-password")
	helper.ReadFromRequestBody(request, &userRequest)
	controller.UserService.UpdatePasswordByUsername(request.Context(), userRequest, username)

	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "password updated",
	}

	helper.TextSuccess("200 OK", "PUT /api/v1/users/current/change-password")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("Authorization")
	helper.TextSuccess("Request", "PUT /api/v1/users/logout")

	controller.UserService.Logout(request.Context(), token)
	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "logged out",
	}

	helper.TextSuccess("200 OK", "PUT /api/v1/users/logout")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller UserControllerImpl) Verification(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserVerificationRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	helper.TextSuccess("Request", "POST /api/v1/users/verification")
	email := controller.UserService.Verification(request.Context(), userRequest)
	request.Header.Set("email", email)

	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "otp sent check your email",
	}

	helper.TextSuccess("200 OK", "POST /api/v1/users/verification")
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller UserControllerImpl) NewPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserNewPasswordRequest{}
	helper.ReadFromRequestBody(request, &userRequest)
	email := request.Header.Get("email")
	helper.TextSuccess("Request", "POST /api/v1/users/new-password")

	err := controller.UserService.NewPassword(request.Context(), email, userRequest)

	if err != nil {
		webResponse := web.ResponseError{
			Code:   400,
			Status: "BAD REQUEST",
			Error:  err.Error(),
		}
		helper.TextFailed("400 BAD REQUEST", "POST /api/v1/users/new-password")
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "success",
	}

	helper.TextSuccess("200 OK", "POST /api/v1/users/new-password")
	helper.WriteToResponseBody(writer, webResponse)
}
