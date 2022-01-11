package repository

import (
	"Go-FoodMarket/model/domain"
	"context"
	"database/sql"
)

type MarketRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Market
	GetRandom(ctx context.Context, tx *sql.Tx) []domain.Market
	GetById(ctx context.Context, tx *sql.Tx, marketId int) (domain.Market, error)
	GetByCategory(ctx context.Context, tx *sql.Tx, categoryId int) []domain.Market
	GetBySearch(ctx context.Context, tx *sql.Tx, searchValue string) []domain.Market
}