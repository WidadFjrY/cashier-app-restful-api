package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/web"
	services "github.com/widadfjry/cashier-app/services/product"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return &ProductControllerImpl{service: service}
}

func (controller ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productRequest)
	productRequest.AddedBy = request.Header.Get("username")

	productResponse := controller.service.Save(request.Context(), productRequest)

	writer.WriteHeader(201)
	webResponse := web.Response{
		Code:   201,
		Status: "CREATED",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) GetProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	productIdInt, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	product := controller.service.FindById(request.Context(), productIdInt)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   product,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) GetProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	page := params.ByName("page")
	pageInt, err := strconv.Atoi(page)
	helper.PanicIfError(err)

	products, pages := controller.service.GetAll(request.Context(), pageInt)
	productsWithPages := web.ProductGetResponseWithPages{
		Items: products,
		Pages: pages,
	}

	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   productsWithPages,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) UpdateProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)
	var productRequest web.ProductUpdateRequest
	helper.ReadFromRequestBody(request, &productRequest)

	modifiedBy := request.Header.Get("username")
	productRequest.ModifiedBy = modifiedBy

	product := controller.service.UpdateById(request.Context(), productRequest, productId)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   product,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller ProductControllerImpl) DeleteProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	controller.service.DeleteById(request.Context(), productId)
	webResponse := web.ResponseMessage{
		Code:    200,
		Status:  "OK",
		Message: "product deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
