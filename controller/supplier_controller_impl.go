package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/suppliers/domain/service"
	"golang-point-of-sales-system/modules/suppliers/dto/request"
	"golang-point-of-sales-system/modules/suppliers/dto/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type SupplierControllerImpl struct {
	SupplierService service.SupplierService
}

func NewSupplierController(SupplierService service.SupplierService) ProductController {
	return &SupplierControllerImpl{
		SupplierService: SupplierService,
	}
}

func (controller *SupplierControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	supplierRequest := request.SupplierCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &supplierRequest)

	productResponse := controller.SupplierService.Create(cRequest.Context(), supplierRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productRequest := request.SupplierUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &productRequest)

	supplierId := params.ByName("supplierId")

	parsedUUID, err := uuid.Parse(supplierId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	productRequest.Id = parsedUUID

	productResponse := controller.SupplierService.Update(cRequest.Context(), productRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	supplierId := params.ByName("supplierId")

	id, err := uuid.Parse(supplierId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.SupplierService.Delete(cRequest.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	supplierId := params.ByName("supplierId")

	id, err := uuid.Parse(supplierId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	productResponse := controller.SupplierService.FindById(cRequest.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SupplierControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productResponses := controller.SupplierService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
