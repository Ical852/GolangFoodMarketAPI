package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/favourite"
)

func ToFavouriteResponse(favouriteDomain domain.Favourite) favourite.FavouriteResponse {
	return favourite.FavouriteResponse{
		Id:       favouriteDomain.Id,
		MarketId: favouriteDomain.MarketId,
		UserId:   favouriteDomain.UserId,
	}
}

func ToFavouriteResponses(favourites []domain.Favourite) []favourite.FavouriteResponse {
	var favouriteResponses []favourite.FavouriteResponse
	for _, favourite := range favourites {
		favouriteResponses = append(favouriteResponses, ToFavouriteResponse(favourite))
	}
	return favouriteResponses
}