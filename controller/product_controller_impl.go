package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/domain/service"
	"golang-point-of-sales-system/modules/products/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productRequest := request.ProductCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &productRequest)

	productResponse := controller.ProductService.Create(cRequest.Context(), productRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productRequest := request.ProductUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &productRequest)

	productId := params.ByName("productId")

	parsedUUID, err := uuid.Parse(productId)
	if err != nil {
		http.Error(writer, "Invalid UUID format", http.StatusBadRequest)
		return
	}
	productRequest.Id = parsedUUID

	productResponse := controller.ProductService.Update(cRequest.Context(), productRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	id, err := uuid.Parse(productId)
	if err != nil {
		http.Error(writer, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	controller.ProductService.Delete(cRequest.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	id, err := uuid.Parse(productId)
	if err != nil {
		http.Error(writer, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	productResponse := controller.ProductService.FindById(cRequest.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	productResponses := controller.ProductService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
