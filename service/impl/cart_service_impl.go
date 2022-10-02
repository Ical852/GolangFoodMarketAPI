package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/helper/response"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/cart"
	"Go-FoodMarket/model/web/transaction"
	"Go-FoodMarket/repository/repository"
	"Go-FoodMarket/service/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CartServiceImpl struct {
	CartRepository repository.CartReposiory
	DB *sql.DB
	Validate *validator.Validate
}

func NewCartService(cartReposiory repository.CartReposiory, DB *sql.DB, validate *validator.Validate) service.CartService {
	return &CartServiceImpl{
		CartRepository: cartReposiory,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *CartServiceImpl) Get(ctx context.Context, userId int) []cart.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	carts := service.CartRepository.Get(ctx, tx, userId)

	return response.ToCartResponses(carts)
}

func (service *CartServiceImpl) Add(ctx context.Context, request cart.CartAddRequest) cart.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		ProductId: request.ProductId,
		UserId: request.UserId,
		Qty: request.Qty,
		Status: request.Status,
	}

	cart = service.CartRepository.Add(ctx, tx, cart)
	return response.ToCartResponse(cart)
}

func (service *CartServiceImpl) Delete(ctx context.Context, request cart.CartDeleteRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		ProductId: request.ProductId,
		UserId: request.UserId,
	}

	service.CartRepository.Delete(ctx, tx, cart)
}

func (service *CartServiceImpl) Transaction(ctx context.Context, request transaction.TransactionCreateRequest) transaction.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaction := domain.Transaction{
		OrderId: request.OrderId,
		GrossAmount: request.GrossAmount,
		UserId: request.UserId,
		ProductId: request.ProductId,
		Status: request.Status,
	}


	user := domain.User{
		FullName: request.FullName,
		Email: request.Email,
		PhoneNumber: request.Phone,
	}

	transaction = service.CartRepository.Transaction(ctx, tx, transaction, user)
	return response.ToTransactionResponse(transaction)
}

func (service *CartServiceImpl) GetTransaction(ctx context.Context, userId int) []transaction.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transactions := service.CartRepository.GetTransaction(ctx, tx, userId)

	return response.ToTransactionResponses(transactions)
}

