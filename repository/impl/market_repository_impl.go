package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/repository/repository"
	"context"
	"database/sql"
	"errors"
)

type MarketRepositoryImpl struct {

}

func NewMarketRepository() repository.MarketRepository {
	return &MarketRepositoryImpl{}
}

func (repository *MarketRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Market {
	SQL := "SELECT * FROM market"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var markets []domain.Market
	for rows.Next() {
		market := domain.Market{}
		err := rows.Scan(&market.Id, &market.CategoryId, &market.MarketName, &market.MarketImg, &market.MarketLocation, &market.MarketRating, &market.MarketProductName, &market.MarketEstimate)
		helper.PanicIfError(err)
		markets = append(markets, market)
	}
	return markets
}

func (repository *MarketRepositoryImpl) GetRandom(ctx context.Context, tx *sql.Tx) []domain.Market {
	SQL := "SELECT * FROM market ORDER BY RAND() LIMIT 5"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var markets []domain.Market
	for rows.Next() {
		market := domain.Market{}
		err := rows.Scan(&market.Id, &market.CategoryId, &market.MarketName, &market.MarketImg, &market.MarketLocation, &market.MarketRating, &market.MarketProductName, &market.MarketEstimate)
		helper.PanicIfError(err)
		markets = append(markets, market)
	}
	return markets
}

func (repository *MarketRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, marketId int) (domain.Market, error) {
	SQL := "SELECT * FROM market WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, marketId)
	helper.PanicIfError(err)
	defer rows.Close()

	market := domain.Market{}
	if rows.Next() {
		err := rows.Scan(&market.Id, &market.CategoryId, &market.MarketName, &market.MarketImg, &market.MarketLocation, &market.MarketRating, &market.MarketProductName, &market.MarketEstimate)
		helper.PanicIfError(err)
		return market, nil
	} else {
		return market, errors.New("Market Data is Not Found")
	}
}

func (repository *MarketRepositoryImpl) GetByCategory(ctx context.Context, tx *sql.Tx, categoryId int) []domain.Market {
	SQL := "SELECT * FROM market WHERE category_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var markets []domain.Market
	for rows.Next() {
		market := domain.Market{}
		err := rows.Scan(&market.Id, &market.CategoryId, &market.MarketName, &market.MarketImg, &market.MarketLocation, &market.MarketRating, &market.MarketProductName, &market.MarketEstimate)
		helper.PanicIfError(err)
		markets = append(markets, market)
	}
	return markets
}

func (repository *MarketRepositoryImpl) GetBySearch(ctx context.Context, tx *sql.Tx, searchValue string) []domain.Market {
	SQL := "SELECT * FROM market WHERE market_name LIKE '"+ "%"+searchValue+"%"+"' "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var markets []domain.Market
	for rows.Next() {
		market := domain.Market{}
		err := rows.Scan(&market.Id, &market.CategoryId, &market.MarketName, &market.MarketImg, &market.MarketLocation, &market.MarketRating, &market.MarketProductName, &market.MarketEstimate)
		helper.PanicIfError(err)
		markets = append(markets, market)
	}
	return markets
}