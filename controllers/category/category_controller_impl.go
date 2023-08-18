package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/web"
	services "github.com/widadfjry/cashier-app/services/category"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &CategoryControllerImpl{service: service}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryRequest)
	username := request.Header.Get("username")
	categoryRequest.AddedBy = username

	categoryResponse := controller.service.Save(request.Context(), categoryRequest)

	writer.WriteHeader(201)
	webResponse := web.Response{
		Code:   201,
		Status: "CREATED",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) GetCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	category := controller.service.FindById(request.Context(), categoryIdInt)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   category,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) GetAllCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := controller.service.FindAll(request.Context())

	webResponse := web.ResponseItems{
		Code:   200,
		Status: "OK",
		Items:  categories,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) UpdateById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryRequest)

	username := request.Header.Get("username")
	categoryRequest.AddedBy = username

	category := controller.service.UpdateById(request.Context(), categoryRequest, categoryIdInt)

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.service.DeleteById(request.Context(), categoryIdInt)

	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "success deleted category",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
