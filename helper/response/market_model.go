package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/market"
)

func ToMarketResponse(marketDomain domain.Market) market.MarketResponse {
	return market.MarketResponse{
		Id:                marketDomain.Id,
		CategoryId:        marketDomain.CategoryId,
		MarketName:        marketDomain.MarketName,
		MarketImg:         marketDomain.MarketImg,
		MarketLocation:    marketDomain.MarketLocation,
		MarketRating:      marketDomain.MarketRating,
		MarketProductName: marketDomain.MarketProductName,
		MarketEstimate:    marketDomain.MarketEstimate,
	}
}

func ToMarketResponses(markets []domain.Market) []market.MarketResponse {
	var marketResponses []market.MarketResponse
	for _, market := range markets {
		marketResponses = append(marketResponses, ToMarketResponse(market))
	}
	return marketResponses
}