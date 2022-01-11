package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MarketController interface {
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetRandom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetByCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetBySearch(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}