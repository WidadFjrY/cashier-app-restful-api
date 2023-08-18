package repositories

import (
	"context"
	"database/sql"
	"github.com/widadfjry/cashier-app/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, categories domain.Categories) domain.CategoriesUpdateOrCreate
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Categories
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Categories
	UpdateById(ctx context.Context, tx *sql.Tx, categories domain.Categories, categoryId int) domain.CategoriesUpdateOrCreate
	DeleteById(ctx context.Context, tx *sql.Tx, categoryId int)
	FindByIdProduct(ctx context.Context, tx *sql.Tx, productId int) []domain.Categories
}
