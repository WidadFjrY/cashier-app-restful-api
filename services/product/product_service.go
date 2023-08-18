package services

import (
	"context"
	"github.com/widadfjry/cashier-app/model/web"
)

type ProductService interface {
	Save(ctx context.Context, request web.ProductCreateRequest) web.ProductCreateResponse
	FindById(ctx context.Context, productId int) web.ProductGetResponse
	GetAll(ctx context.Context, page int) ([]web.ProductGetResponse, web.Pages)
	UpdateById(ctx context.Context, request web.ProductUpdateRequest, productId int) web.ProductUpdateResponse
	DeleteById(ctx context.Context, productId int)
}
