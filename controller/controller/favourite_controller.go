package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type FavouriteController interface {
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Check(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Add(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
