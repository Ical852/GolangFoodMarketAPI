package impl

import (
	"Go-FoodMarket/exception"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/helper/response"
	"Go-FoodMarket/model/web/market"
	"Go-FoodMarket/repository/repository"
	"Go-FoodMarket/service/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type MarketServiceImpl struct {
	MarketRepository repository.MarketRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewMarketService(marketRepository repository.MarketRepository, DB *sql.DB, validate *validator.Validate) service.MarketService {
	return &MarketServiceImpl{
		MarketRepository: marketRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *MarketServiceImpl) Get(ctx context.Context) []market.MarketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	markets := service.MarketRepository.Get(ctx, tx)
	return response.ToMarketResponses(markets)
}

func (service *MarketServiceImpl) GetRandom(ctx context.Context) []market.MarketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	markets := service.MarketRepository.GetRandom(ctx, tx)
	return response.ToMarketResponses(markets)
}

func (service *MarketServiceImpl) GetById(ctx context.Context, marketId int) market.MarketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	market, err := service.MarketRepository.GetById(ctx, tx, marketId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return response.ToMarketResponse(market)
}

func (service *MarketServiceImpl) GetByCategory(ctx context.Context, categoryId int) []market.MarketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	markets := service.MarketRepository.GetByCategory(ctx, tx, categoryId)
	return response.ToMarketResponses(markets)
}

func (service *MarketServiceImpl) GetBySearch(ctx context.Context, searchValue string) []market.MarketResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	markets := service.MarketRepository.GetBySearch(ctx, tx, searchValue)
	return response.ToMarketResponses(markets)
}
