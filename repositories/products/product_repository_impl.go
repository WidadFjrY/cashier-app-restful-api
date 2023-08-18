package repositories

import (
	"context"
	"database/sql"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/domain"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, products domain.Products) domain.ProductCreateOrUpdate {
	SQL := "INSERT INTO products (sku, name, description, price, stock, brand, weight, dimension, variant, added_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, products.SKU, products.Name, products.Description, products.Price, products.Stock, products.Brand, products.Weight, products.Dimension, products.Variant, products.AddedBy)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	return domain.ProductCreateOrUpdate{
		Id:          int(id),
		SKU:         products.SKU,
		Name:        products.Name,
		Description: products.Description,
		Price:       products.Price,
		Stock:       products.Stock,
		Brand:       products.Brand,
		Weight:      products.Weight,
		Dimension:   products.Dimension,
		Variant:     products.Variant,
		AddedBy:     products.AddedBy,
	}
}

func (repository ProductRepositoryImpl) IfSKUExist(ctx context.Context, tx *sql.Tx, sku string) {
	SQL := "SELECT sku FROM products WHERE sku = ?"
	rows, err := tx.QueryContext(ctx, SQL, sku)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		panic(exception.NewBadRequestErrors("sku already exist"))
	}
}

func (repository ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) domain.Products {
	SQL := "SELECT * FROM products WHERE id = ?"
	rowsProducts, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rowsProducts.Close()

	var product domain.Products
	if rowsProducts.Next() {
		err := rowsProducts.Scan(
			&product.Id,
			&product.SKU,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Brand,
			&product.Weight,
			&product.Dimension,
			&product.Variant,
			&product.AddedBy,
			&product.ModifiedBy,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		helper.PanicIfError(err)
	} else {
		panic(exception.NewNotFoundErrors("product not found"))
	}

	return product
}

func (repository ProductRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx, offset int) []domain.ProductJoinWithCategory {
	SQL := "SELECT * FROM products JOIN categories category on products.id = category.product_id LIMIT 10 OFFSET ?"
	rows, err := tx.QueryContext(ctx, SQL, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.ProductJoinWithCategory
	for rows.Next() {
		var product domain.ProductJoinWithCategory
		rows.Scan(
			&product.Id,
			&product.SKU,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.Brand,
			&product.Weight,
			&product.Dimension,
			&product.Variant,
			&product.AddedBy,
			&product.ModifiedBy,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.IdCategory,
			&product.NameCategory,
			&product.CategoryAddedBy,
			&product.CategoryModifiedBy,
			&product.CategoryCreatedAt,
			&product.CategoryUpdatedAt,
		)

		products = append(products, product)
	}

	return products
}

func (repository ProductRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(*) AS total_product FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var totalProduct int
	if rows.Next() {
		err := rows.Scan(&totalProduct)
		helper.PanicIfError(err)
	}

	return totalProduct
}

func (repository ProductRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, products domain.Products, productId int) domain.ProductCreateOrUpdate {
	SQL := "UPDATE products SET sku = ?, name = ?, description = ?, price = ?, stock = ?, brand = ?, weight = ?, dimension = ?, variant = ?, modified_by = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, products.SKU, products.Name, products.Description, products.Price, products.Stock, products.Brand, products.Weight, products.Dimension, products.Variant, products.ModifiedBy, productId)
	helper.PanicIfError(err)

	return domain.ProductCreateOrUpdate{
		Id:          productId,
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
	}
}

func (repository ProductRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, productId int) {
	SQL := "DELETE FROM products WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(err)
}

func (repository ProductRepositoryImpl) GetId(ctx context.Context, tx *sql.Tx) []int {
	SQL := "SELECT id FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		helper.PanicIfError(err)

		ids = append(ids, id)
	}

	return ids
}
