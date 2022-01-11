package impl

import (
	"Go-FoodMarket/exception"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/helper/response"
	"Go-FoodMarket/model/web/products"
	"Go-FoodMarket/repository/repository"
	"Go-FoodMarket/service/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ProductsServiceImpl struct {
	ProductRepository repository.ProductsRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewProductsService(productsRepository repository.ProductsRepository, DB *sql.DB, validate *validator.Validate) service.ProductsService  {
	return &ProductsServiceImpl{
		ProductRepository: productsRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductsServiceImpl) Get(ctx context.Context, marketId int) []products.ProductsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.Get(ctx, tx, marketId)
	return response.ToProductsResponses(products)
}

func (service *ProductsServiceImpl) GetById(ctx context.Context, productId int) products.ProductsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.GetById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return response.ToProductsResponse(product)
}