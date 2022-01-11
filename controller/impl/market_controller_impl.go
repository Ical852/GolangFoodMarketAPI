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

type MarketControllerImpl struct {
	MarketService service.MarketService
}

func NewMarketController(marketService service.MarketService) controller.MarketController {
	return &MarketControllerImpl{
		MarketService: marketService,
	}
}

func (controller *MarketControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	marketResponse := controller.MarketService.Get(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   marketResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MarketControllerImpl) GetRandom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	marketResponse := controller.MarketService.GetRandom(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   marketResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MarketControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	marketId := params.ByName("marketId")
	id, err := strconv.Atoi(marketId)
	helper.PanicIfError(err)

	marketResponse := controller.MarketService.GetById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   marketResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MarketControllerImpl) GetByCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categorytId := params.ByName("categoryId")
	id, err := strconv.Atoi(categorytId)
	helper.PanicIfError(err)

	marketResponse := controller.MarketService.GetByCategory(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   marketResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MarketControllerImpl) GetBySearch(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	searchValue := params.ByName("searchValue")
	search := string(searchValue)

	marketResponse := controller.MarketService.GetBySearch(request.Context(), search)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   marketResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
