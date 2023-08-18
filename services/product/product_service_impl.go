package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/domain"
	"github.com/widadfjry/cashier-app/model/web"
	repositories2 "github.com/widadfjry/cashier-app/repositories/categories"
	repositories "github.com/widadfjry/cashier-app/repositories/products"
	"math"
)

type ProductServiceImpl struct {
	ProductRepository  repositories.ProductRepository
	CategoryRepository repositories2.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewProductService(productRepository repositories.ProductRepository, categoryRepository repositories2.CategoryRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository, CategoryRepository: categoryRepository, DB: DB, Validate: validate}
}

func (service ProductServiceImpl) Save(ctx context.Context, request web.ProductCreateRequest) web.ProductCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	service.ProductRepository.IfSKUExist(ctx, tx, request.SKU)

	productRequest := domain.Products{
		SKU:         request.SKU,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Brand:       request.Brand,
		Weight:      request.Weight,
		Dimension:   request.Dimension,
		Variant:     request.Variant,
		AddedBy:     request.AddedBy,
	}
	product := service.ProductRepository.Save(ctx, tx, productRequest)
	return helper.ToProductCreateResponse(product)
}

func (service ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductGetResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tx2, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx2)

	product := service.ProductRepository.FindById(ctx, tx, productId)
	categories := service.CategoryRepository.FindByIdProduct(ctx, tx, productId)

	var webCategories []web.Categories

	for _, category := range categories {
		webCategory := web.Categories{
			IdCategory: category.Id,
			DataCategory: web.GetDataCategory{
				ProductId:  category.ProductId,
				Name:       category.Name,
				AddedBy:    category.AddedBy,
				ModifiedBy: category.ModifiedBy,
				CreatedAt:  category.CreatedAt,
				UpdatedAt:  category.UpdatedAt,
			},
		}
		webCategories = append(webCategories, webCategory)
	}

	return helper.ToProductResponse(product, webCategories)
}

func (service ProductServiceImpl) GetAll(ctx context.Context, page int) ([]web.ProductGetResponse, web.Pages) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	totalProduct := service.ProductRepository.Count(ctx, tx)
	totalPages := math.Ceil(float64(totalProduct) / 10.00)
	pageSize := 10

	if float64(page) > totalPages {
		panic(exception.NewNotFoundErrors("page not found"))
	}

	startIndex := (page - 1) * pageSize
	endIndex := page * pageSize

	if endIndex > totalProduct {
		endIndex = totalProduct
	}

	pages := web.Pages{
		Page:       page,
		TotalPages: totalPages,
		TotalItems: totalProduct,
	}

	var webResponses []web.ProductGetResponse

	ids := service.ProductRepository.GetId(ctx, tx)
	for _, id := range ids {
		var categoriesResponse []web.Categories
		product := service.ProductRepository.FindById(ctx, tx, id)
		categories := service.CategoryRepository.FindByIdProduct(ctx, tx, id)

		for _, category := range categories {
			webCategory := web.Categories{
				IdCategory: category.Id,
				DataCategory: web.GetDataCategory{
					ProductId:  category.ProductId,
					Name:       category.Name,
					AddedBy:    category.AddedBy,
					ModifiedBy: category.ModifiedBy,
					CreatedAt:  category.CreatedAt,
					UpdatedAt:  category.UpdatedAt,
				},
			}
			categoriesResponse = append(categoriesResponse, webCategory)
		}
		webResponse := helper.ToProductResponse(product, categoriesResponse)
		webResponses = append(webResponses, webResponse)
	}

	webResponseWithOffset := webResponses[startIndex:endIndex]
	return webResponseWithOffset, pages

}

func (service ProductServiceImpl) UpdateById(ctx context.Context, request web.ProductUpdateRequest, productId int) web.ProductUpdateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	helper.PanicIfError(err)

	productRequest := domain.Products{
		SKU:         request.SKU,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Brand:       request.Brand,
		Weight:      request.Weight,
		Dimension:   request.Dimension,
		Variant:     request.Variant,
		ModifiedBy:  request.ModifiedBy,
	}

	service.CategoryRepository.FindById(ctx, tx, request.IdCategory)
	product := service.ProductRepository.UpdateById(ctx, tx, productRequest, productId)
	return helper.ToProductUpdateResponse(product)
}

func (service ProductServiceImpl) DeleteById(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//service.ProductRepository.FindById(ctx, tx, productId)
	service.ProductRepository.DeleteById(ctx, tx, productId)
}
