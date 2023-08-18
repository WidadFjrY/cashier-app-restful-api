package repositories

import (
	"context"
	"database/sql"
	"github.com/widadfjry/cashier-app/exception"
	"github.com/widadfjry/cashier-app/helper"
	"github.com/widadfjry/cashier-app/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, categories domain.Categories) domain.CategoriesUpdateOrCreate {
	SQL := "INSERT INTO categories (product_id, name, added_by, modified_by) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, categories.ProductId, categories.Name, categories.AddedBy, categories.ModifiedBy)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	categoryCreate := domain.CategoriesUpdateOrCreate{
		Id:        int(id),
		ProductId: categories.ProductId,
		Name:      categories.Name,
		AddedBy:   categories.AddedBy,
	}

	return categoryCreate
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Categories {
	SQL := "SELECT * FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var category domain.Categories
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.ProductId, &category.Name, &category.AddedBy, &category.ModifiedBy, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)
	} else {
		panic(exception.NewNotFoundErrors("category not found"))
	}

	return category
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Categories {
	SQL := "SELECT * FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Categories
	for rows.Next() {
		var category domain.Categories
		err := rows.Scan(&category.Id, &category.ProductId, &category.Name, &category.AddedBy, &category.ModifiedBy, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, categories domain.Categories, categoryId int) domain.CategoriesUpdateOrCreate {
	SQL := "UPDATE categories SET product_id = ?, name = ?, modified_by = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categories.ProductId, categories.Name, categories.AddedBy, categoryId)
	helper.PanicIfError(err)

	category := domain.CategoriesUpdateOrCreate{
		Id:        categoryId,
		ProductId: categories.ProductId,
		Name:      categories.Name,
		AddedBy:   categories.AddedBy,
	}

	return category
}

func (repository *CategoryRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, categoryId int) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindByIdProduct(ctx context.Context, tx *sql.Tx, productId int) []domain.Categories {
	SQL := "SELECT * FROM categories WHERE product_id = ?"
	rowsCategories, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rowsCategories.Close()

	var categories []domain.Categories
	for rowsCategories.Next() {
		var category domain.Categories
		err := rowsCategories.Scan(&category.Id, &category.ProductId, &category.Name, &category.AddedBy, &category.ModifiedBy, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
