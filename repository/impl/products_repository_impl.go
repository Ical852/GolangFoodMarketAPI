package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/repository/repository"
	"context"
	"database/sql"
	"errors"
)

type ProductsRepositoryImpl struct {
	
}

func NewProductRepository() repository.ProductsRepository {
	return &ProductsRepositoryImpl{}
}

func (repository *ProductsRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, marketId int) []domain.Products {
	SQL := "SELECT * FROM products WHERE market_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, marketId)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Products
	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product.Id, &product.MarketId, &product.ProductName, &product.ProductPrice, &product.ProductImg, &product.ProductStock, &product.ProductStatus)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (repository *ProductsRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, productId int) (domain.Products, error) {
	SQL := "SELECT * FROM products WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	products := domain.Products{}
	if rows.Next() {
		err := rows.Scan(&products.Id, &products.MarketId, &products.ProductName, &products.ProductPrice, &products.ProductImg, &products.ProductStock, &products.ProductStatus)
		helper.PanicIfError(err)
		return products, nil
	} else {
		return products, errors.New("Product Tidak Ditemukan")
	}
}