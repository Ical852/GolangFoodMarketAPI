package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/products"
)

func ToProductsResponse(productsDomain domain.Products) products.ProductsResponse {
	return products.ProductsResponse{
		Id:            productsDomain.Id,
		MarketId:      productsDomain.MarketId,
		ProductName:   productsDomain.ProductName,
		ProductPrice:  productsDomain.ProductPrice,
		ProductImg:    productsDomain.ProductImg,
		ProductStock:  productsDomain.ProductStock,
		ProductStatus: productsDomain.ProductStatus,
	}
}

func ToProductsResponses(productsDomain []domain.Products) []products.ProductsResponse {
	var productsResponses []products.ProductsResponse
	for _, products := range productsDomain {
		productsResponses = append(productsResponses, ToProductsResponse(products))
	}
	return productsResponses
}