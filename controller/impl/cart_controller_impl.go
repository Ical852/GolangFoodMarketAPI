package impl

import (
	"Go-FoodMarket/controller/controller"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/web"
	"Go-FoodMarket/model/web/cart"
	"Go-FoodMarket/service/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CartControllerImpl struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) controller.CartController {
	return &CartControllerImpl{
		CartService: cartService,
	}
}

func (controller *CartControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	cartResponse := controller.CartService.Get(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) Add(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartAddRequest := cart.CartAddRequest{}
	helper.ReadFromRequestBody(request, &cartAddRequest)

	cartResponse := controller.CartService.Add(request.Context(), cartAddRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartDeleteRequest := cart.CartDeleteRequest{}
	helper.ReadFromRequestBody(request, &cartDeleteRequest)

	controller.CartService.Delete(request.Context(), cartDeleteRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
