package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/penjualan/domain/service"
	"golang-point-of-sales-system/modules/penjualan/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type PenjualanControllerImpl struct {
	PenjualanService service.PenjualanService
}

func NewPenjualanController(penjualanService service.PenjualanService) PenjualanController {
	return &PenjualanControllerImpl{
		PenjualanService: penjualanService,
	}
}

func (controller *PenjualanControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataRequest := request.PenjualanCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &dataRequest)

	dataResponse := controller.PenjualanService.Create(cRequest.Context(), dataRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PenjualanControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataRequest := request.PenjualanUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &dataRequest)

	penjualanId := params.ByName("penjualanId")

	parsedUUID, err := uuid.Parse(penjualanId)
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

	dataResponse := controller.PenjualanService.Update(cRequest.Context(), dataRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PenjualanControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	penjualanId := params.ByName("penjualanId")

	parsedUUID, err := uuid.Parse(penjualanId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.PenjualanService.Delete(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Data deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PenjualanControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	penjualanId := params.ByName("penjualanId")

	parsedUUID, err := uuid.Parse(penjualanId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	dataResponse := controller.PenjualanService.FindById(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PenjualanControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataResponse := controller.PenjualanService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
