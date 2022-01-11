package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductsController interface {
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}