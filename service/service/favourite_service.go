package service

import (
	"Go-FoodMarket/model/web/favourite"
	"context"
)

type FavouriteService interface {
	Get(ctx context.Context, userId int) []favourite.FavouriteResponse
	Check(ctx context.Context, request favourite.FavouriteCheckRequest) favourite.FavouriteResponse
	Add(ctx context.Context, request favourite.FavouriteAddRequest) favourite.FavouriteResponse
	Delete(ctx context.Context, request favourite.FavouriteDeleteRequest)
}