package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/pembelianDetail/domain/service"
	"golang-point-of-sales-system/modules/pembelianDetail/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type PembelianDetailControllerImpl struct {
	PembelianDetailService service.PembelianDetailService
}

func NewPembelianDetailController(PembelianDetailController service.PembelianDetailService) PembelianDetailController {
	return &PembelianDetailControllerImpl{
		PembelianDetailService: PembelianDetailController,
	}
}

func (controller *PembelianDetailControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataRequest := request.PembelianDetailCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &dataRequest)

	dataResponse := controller.PembelianDetailService.Create(cRequest.Context(), dataRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PembelianDetailControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataRequest := request.PembelianDetailUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &dataRequest)

	pembelianDetailId := params.ByName("pembelianDetailId")

	parsedUUID, err := uuid.Parse(pembelianDetailId)
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

	dataResponse := controller.PembelianDetailService.Update(cRequest.Context(), dataRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PembelianDetailControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	pembelianDetailId := params.ByName("pembelianDetailId")

	parsedUUID, err := uuid.Parse(pembelianDetailId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.PembelianDetailService.Delete(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Data deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *PembelianDetailControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	pembelianDetailId := params.ByName("pembelianDetailId")

	parsedUUID, err := uuid.Parse(pembelianDetailId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	dataResponse := controller.PembelianDetailService.FindById(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PembelianDetailControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	dataResponse := controller.PembelianDetailService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
