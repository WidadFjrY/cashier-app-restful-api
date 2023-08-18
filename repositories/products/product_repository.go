package repositories

import (
	"context"
	"database/sql"
	"github.com/widadfjry/cashier-app/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, products domain.Products) domain.ProductCreateOrUpdate
	IfSKUExist(ctx context.Context, tx *sql.Tx, sku string)
	FindById(ctx context.Context, tx *sql.Tx, productId int) domain.Products
	GetAll(ctx context.Context, tx *sql.Tx, offset int) []domain.ProductJoinWithCategory
	Count(ctx context.Context, tx *sql.Tx) int
	UpdateById(ctx context.Context, tx *sql.Tx, products domain.Products, productId int) domain.ProductCreateOrUpdate
	DeleteById(ctx context.Context, tx *sql.Tx, productId int)
	GetId(ctx context.Context, tx *sql.Tx) []int
}
