package service

import (
	"Go-FoodMarket/model/web/cart"
	"context"
)

type CartService interface {
	Get(ctx context.Context, userId int) []cart.CartResponse
	Add(ctx context.Context, request cart.CartAddRequest) cart.CartResponse
	Delete(ctx context.Context, request cart.CartDeleteRequest)
}