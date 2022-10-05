package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/repository/repository"
	"context"
	"database/sql"
	"github.com/veritrans/go-midtrans"
	"strconv"
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

func (repository *CartRepositoryImpl) Transaction(ctx context.Context, tx *sql.Tx, transaction domain.Transaction, user domain.User) domain.Transaction {
	paymentUrl := ""
	grossAmount := 0
	if transaction.Status == "nolink" {
		paymentUrl = "join"
		grossAmount = 0
	} else {
		grossAmount = transaction.GrossAmount

		midclient := midtrans.NewClient()
		midclient.APIEnvType = midtrans.Sandbox

		snapGateway := midtrans.SnapGateway{
			Client: midclient,
		}

		snapReq := &midtrans.SnapReq{
			CustomerDetail: &midtrans.CustDetail{
				FName:    user.FullName,
				Email:    user.Email,
				Phone: user.PhoneNumber,
			},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(transaction.OrderId),
				GrossAmt: int64(transaction.GrossAmount),
			},
		}

		snapTokenResp, _ := snapGateway.GetToken(snapReq)
		paymentUrl = snapTokenResp.RedirectURL
	}

	SQL := "INSERT INTO transaction(order_id, gross_amount, payment_url, user_id, product_id, status) VALUES(?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, transaction.OrderId, grossAmount, paymentUrl, transaction.UserId, transaction.ProductId, transaction.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	transaction.Id = int(id)
	transaction.PaymentUrl = paymentUrl

	return transaction
}

func (repository *CartRepositoryImpl) GetTransaction(ctx context.Context, tx *sql.Tx, userId int) []domain.Transaction {
	SQL := "SELECT * FROM transaction WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.OrderId, &transaction.GrossAmount,  &transaction.PaymentUrl, &transaction.UserId, &transaction.ProductId, &transaction.Status)
		helper.PanicIfError(err)
		transactions = append(transactions, transaction)
	}
	return transactions
}

