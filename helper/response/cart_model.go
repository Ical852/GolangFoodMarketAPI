package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/cart"
)

func ToCartResponse(cartDomain domain.Cart) cart.CartResponse {
	return cart.CartResponse{
		Id:        cartDomain.Id,
		ProductId: cartDomain.ProductId,
		UserId:    cartDomain.UserId,
		Qty:       cartDomain.Qty,
		Status:    cartDomain.Status,
	}
}

func ToCartResponses(carts []domain.Cart) []cart.CartResponse {
	var cartResponses []cart.CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, ToCartResponse(cart))
	}
	return cartResponses
}