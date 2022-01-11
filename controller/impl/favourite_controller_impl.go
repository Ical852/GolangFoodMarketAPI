package impl

import (
	"Go-FoodMarket/controller/controller"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/web"
	"Go-FoodMarket/model/web/favourite"
	"Go-FoodMarket/service/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type FavouriteControllerImpl struct {
	FavouriteService service.FavouriteService
}

func NewFavouriteController(favouriteService service.FavouriteService) controller.FavouriteController {
	return &FavouriteControllerImpl{
		FavouriteService: favouriteService,
	}
}

func (controller *FavouriteControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	favouriteResponse := controller.FavouriteService.Get(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   favouriteResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FavouriteControllerImpl) Check(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	favouriteCheckRequest := favourite.FavouriteCheckRequest{}
	helper.ReadFromRequestBody(request, &favouriteCheckRequest)

	favouriteResponse := controller.FavouriteService.Check(request.Context(), favouriteCheckRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   favouriteResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FavouriteControllerImpl) Add(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	favouriteAddrequest := favourite.FavouriteAddRequest{}
	helper.ReadFromRequestBody(request, &favouriteAddrequest)

	favouriteResponse := controller.FavouriteService.Add(request.Context(), favouriteAddrequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   favouriteResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FavouriteControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	favouriteDeleteRequest := favourite.FavouriteDeleteRequest{}
	helper.ReadFromRequestBody(request, &favouriteDeleteRequest)

	controller.FavouriteService.Delete(request.Context(), favouriteDeleteRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
