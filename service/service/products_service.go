package service

import (
	"Go-FoodMarket/model/web/products"
	"context"
)

type ProductsService interface {
	Get(ctx context.Context, marketId int) []products.ProductsResponse
	GetById(ctx context.Context, productId int) products.ProductsResponse
}