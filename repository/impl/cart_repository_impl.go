package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/repository/repository"
	"context"
	"database/sql"
)

type CartRepositoryImpl struct {

}

func NewCartRepository() repository.CartReposiory {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, userId int) []domain.Cart {
	SQL := "SELECT * FROM cart WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.ProductId, &cart.UserId, &cart.Qty, &cart.Status)
		helper.PanicIfError(err)
		carts = append(carts, cart)
	}
	return carts
}

func (repository *CartRepositoryImpl) Add(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "INSERT INTO cart(product_id, user_id, qty, status) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, cart.ProductId, cart.UserId, cart.Qty, cart.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cart.Id = int(id)
	return cart
}

func (repository *CartRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, cart domain.Cart) {
	SQL := "DELETE FROM cart WHERE product_id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.ProductId, cart.UserId)
	helper.PanicIfError(err)
}
