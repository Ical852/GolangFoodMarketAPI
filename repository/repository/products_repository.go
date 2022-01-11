package repository

import (
	"Go-FoodMarket/model/domain"
	"context"
	"database/sql"
)

type ProductsRepository interface {
	Get(ctx context.Context, tx *sql.Tx, marketId int) []domain.Products
	GetById(ctx context.Context, tx *sql.Tx, productId int) (domain.Products, error)
}