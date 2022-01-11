package impl

import (
	"Go-FoodMarket/controller/controller"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/web"
	"Go-FoodMarket/service/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) controller.CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.Get(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
