package services

import (
	"context"
	"github.com/widadfjry/cashier-app/model/web"
)

type CategoryService interface {
	Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateOrUpdateResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	UpdateById(ctx context.Context, request web.CategoryUpdateRequest, categoryId int) web.CategoryCreateOrUpdateResponse
	DeleteById(ctx context.Context, categoryId int)
}
