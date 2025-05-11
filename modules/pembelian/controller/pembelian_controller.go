package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PembelianController interface {
	Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params)
}
