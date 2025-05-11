package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/pembelian/domain/service"
	"golang-point-of-sales-system/modules/pembelian/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type PembelianControllerImpl struct {
	PembelianService service.PembelianService
}

func NewPembelianController(PembelianController service.PembelianService) PembelianController {
	return &PembelianControllerImpl{
		PembelianService: PembelianController,
	}
}

func (controller *PembelianControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataRequest := request.PembelianCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &dataRequest)

	dataResponse := controller.PembelianService.Create(cRequest.Context(), dataRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PembelianControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataRequest := request.PembelianUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &dataRequest)

	pembelianId := params.ByName("pembelianId")

	parsedUUID, err := uuid.Parse(pembelianId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	dataRequest.Id = parsedUUID

	dataResponse := controller.PembelianService.Update(cRequest.Context(), dataRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PembelianControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	pembelianId := params.ByName("pembelianId")

	parsedUUID, err := uuid.Parse(pembelianId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.PembelianService.Delete(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Data deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PembelianControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	pembelianId := params.ByName("pembelianId")

	parsedUUID, err := uuid.Parse(pembelianId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	dataResponse := controller.PembelianService.FindById(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PembelianControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataResponse := controller.PembelianService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
