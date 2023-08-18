package helper

import (
	"github.com/widadfjry/cashier-app/model/domain"
	"github.com/widadfjry/cashier-app/model/web"
)

func ToProductCreateResponse(product domain.ProductCreateOrUpdate) web.ProductCreateResponse {
	return web.ProductCreateResponse{
		Id: product.Id,
		Data: web.CreateDataProduct{
			SKU:         product.SKU,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Brand:       product.Brand,
			Weight:      product.Weight,
			Dimension:   product.Dimension,
			Variant:     product.Variant,
			AddedBy:     product.AddedBy,
		},
	}
}

func ToProductUpdateResponse(product domain.ProductCreateOrUpdate) web.ProductUpdateResponse {
	return web.ProductUpdateResponse{
		Id: product.Id,
		Data: web.UpdateDataProduct{
			SKU:         product.SKU,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Brand:       product.Brand,
			Weight:      product.Weight,
			Dimension:   product.Dimension,
			Variant:     product.Variant,
			ModifiedBy:  product.ModifiedBy,
		},
	}
}

func ToProductResponse(products domain.Products, category []web.Categories) web.ProductGetResponse {
	return web.ProductGetResponse{
		IdProduct: products.Id,
		DataProduct: web.GetDataProduct{
			SKU:         products.SKU,
			Name:        products.Name,
			Description: products.Description,
			Price:       products.Price,
			Stock:       products.Stock,
			Brand:       products.Brand,
			Weight:      products.Weight,
			Dimension:   products.Dimension,
			Variant:     products.Variant,
			ModifiedBy:  products.ModifiedBy,
			AddedBy:     products.AddedBy,
			CreatedAt:   products.CreatedAt,
			UpdatedAt:   products.UpdatedAt,
		},
		Categories: category,
	}
}
