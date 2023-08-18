package helper

import (
	"github.com/widadfjry/cashier-app/model/domain"
	"github.com/widadfjry/cashier-app/model/web"
)

func ToCategoryResponse(categories domain.Categories) web.CategoryResponse {
	return web.CategoryResponse{
		Id: categories.Id,
		Data: web.GetDataCategory{
			ProductId:  categories.ProductId,
			Name:       categories.Name,
			AddedBy:    categories.AddedBy,
			ModifiedBy: categories.ModifiedBy,
			CreatedAt:  categories.CreatedAt,
			UpdatedAt:  categories.UpdatedAt,
		},
	}
}

func ToCategoryUpdateOrCreateResponse(category domain.CategoriesUpdateOrCreate) web.CategoryCreateOrUpdateResponse {
	return web.CategoryCreateOrUpdateResponse{
		Id: category.Id,
		Data: web.CreateOrUpdateDataCategory{
			ProductId: category.ProductId,
			Name:      category.Name,
			AddedBy:   category.AddedBy,
		},
	}
}

func ToCategoryResponses(categories []domain.Categories) []web.CategoryResponse {
	var categoriesResponses []web.CategoryResponse
	for _, category := range categories {
		categoriesResponses = append(categoriesResponses, ToCategoryResponse(category))
	}
	return categoriesResponses
}
