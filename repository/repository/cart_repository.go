package repository

import (
	"Go-FoodMarket/model/domain"
	"context"
	"database/sql"
)

type CartReposiory interface {
	Get(ctx context.Context, tx *sql.Tx, userId int) []domain.Cart
	Add(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	Delete(ctx context.Context, tx *sql.Tx, cart domain.Cart)
	Transaction(ctx context.Context, tx *sql.Tx, transaction domain.Transaction, user domain.User) domain.Transaction
	GetTransaction(ctx context.Context, tx *sql.Tx, userId int) []domain.Transaction
}