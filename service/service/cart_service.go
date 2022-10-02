package service

import (
	"Go-FoodMarket/model/web/cart"
	"Go-FoodMarket/model/web/transaction"
	"context"
)

type CartService interface {
	Get(ctx context.Context, userId int) []cart.CartResponse
	Add(ctx context.Context, request cart.CartAddRequest) cart.CartResponse
	Delete(ctx context.Context, request cart.CartDeleteRequest)
	Transaction(ctx context.Context, request transaction.TransactionCreateRequest) transaction.TransactionResponse
	GetTransaction(ctx context.Context, userId int) []transaction.TransactionResponse
}