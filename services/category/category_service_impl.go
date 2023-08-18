package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/domain"
	"github.com/widadfjry/cashier-app/model/web"
	repositories "github.com/widadfjry/cashier-app/repositories/categories"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewCategoryService(categoryRepository repositories.CategoryRepository, DB *sql.DB, validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validator: validator}
}

func (service CategoryServiceImpl) Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateOrUpdateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validator.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	category := domain.Categories{
		ProductId:  request.ProductId,
		Name:       request.Name,
		AddedBy:    request.AddedBy,
		ModifiedBy: request.AddedBy,
	}

	categoryCreate := service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryUpdateOrCreateResponse(categoryCreate)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := service.CategoryRepository.FindById(ctx, tx, categoryId)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}

func (service CategoryServiceImpl) UpdateById(ctx context.Context, request web.CategoryUpdateRequest, categoryId int) web.CategoryCreateOrUpdateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validator.Struct(request)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	categoryRequest := domain.Categories{
		Id:        categoryId,
		ProductId: request.ProductId,
		Name:      request.Name,
		AddedBy:   request.AddedBy,
	}

	category := service.CategoryRepository.UpdateById(ctx, tx, categoryRequest, categoryId)
	return helper.ToCategoryUpdateOrCreateResponse(category)
}

func (service CategoryServiceImpl) DeleteById(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.CategoryRepository.FindById(ctx, tx, categoryId)
	service.CategoryRepository.DeleteById(ctx, tx, categoryId)
}
