package impl

import (
	"Go-FoodMarket/controller/controller"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/web"
	"Go-FoodMarket/service/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ProductsControllerImpl struct {
	ProductsService service.ProductsService
}

func NewProductsController(productsService service.ProductsService) controller.ProductsController {
	return &ProductsControllerImpl{
		ProductsService: productsService,
	}
}

func (controller *ProductsControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	marketId := params.ByName("marketId")
	id, err := strconv.Atoi(marketId)
	helper.PanicIfError(err)

	productsResponse := controller.ProductsService.Get(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductsControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productsResponse := controller.ProductsService.GetById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
