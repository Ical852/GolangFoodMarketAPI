package service

import (
	"Go-FoodMarket/model/web/market"
	"context"
)

type MarketService interface {
	Get(ctx context.Context) []market.MarketResponse
	GetRandom(ctx context.Context) []market.MarketResponse
	GetById(ctx context.Context, marketId int) market.MarketResponse
	GetByCategory(ctx context.Context, categoryId int) []market.MarketResponse
	GetBySearch(ctx context.Context, searchValue string) []market.MarketResponse
}