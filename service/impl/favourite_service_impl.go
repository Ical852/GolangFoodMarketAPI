package impl

import (
	"Go-FoodMarket/exception"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/helper/response"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/favourite"
	"Go-FoodMarket/repository/repository"
	"Go-FoodMarket/service/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type FavouriteServiceImpl struct {
	FavouriteRepository repository.FavouriteRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewFavouriteService(favouriteRepository repository.FavouriteRepository, DB *sql.DB, validate *validator.Validate) service.FavouriteService {
	return &FavouriteServiceImpl{
		FavouriteRepository: favouriteRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *FavouriteServiceImpl) Get(ctx context.Context, userId int) []favourite.FavouriteResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	favourites := service.FavouriteRepository.Get(ctx, tx, userId)
	return response.ToFavouriteResponses(favourites)
}

func (service *FavouriteServiceImpl) Check(ctx context.Context, request favourite.FavouriteCheckRequest) favourite.FavouriteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	favourite := domain.Favourite{
		MarketId: request.MarketId,
		UserId: request.UserId,
	}

	favourite, err = service.FavouriteRepository.Check(ctx, tx, favourite)

	return response.ToFavouriteResponse(favourite)
}

func (service *FavouriteServiceImpl) Add(ctx context.Context, request favourite.FavouriteAddRequest) favourite.FavouriteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	favourite := domain.Favourite{
		MarketId: request.MarketId,
		UserId: request.UserId,
	}

	favourite = service.FavouriteRepository.Add(ctx, tx, favourite)

	return response.ToFavouriteResponse(favourite)
}

func (service *FavouriteServiceImpl) Delete(ctx context.Context, request favourite.FavouriteDeleteRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	favourite := domain.Favourite{
		MarketId: request.MarketId,
		UserId: request.UserId,
	}

	favourite, err = service.FavouriteRepository.Check(ctx, tx, favourite)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.FavouriteRepository.Delete(ctx, tx, favourite)
}
