package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/categories/domain/service"
	"golang-point-of-sales-system/modules/categories/dto/request"
	"golang-point-of-sales-system/pkg/utils/response"

	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(CategoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: CategoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	categoryRequest := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &categoryRequest)

	categoryResponse := controller.CategoryService.Create(cRequest.Context(), categoryRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	categoryRequest := request.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &categoryRequest)

	categoryId := params.ByName("categoryId")

	parsedUUID, err := uuid.Parse(categoryId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	categoryRequest.Id = parsedUUID

	categoryResponse := controller.CategoryService.Update(cRequest.Context(), categoryRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	parsedUUID, err := uuid.Parse(categoryId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.CategoryService.Delete(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Category deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	parsedUUID, err := uuid.Parse(categoryId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	categoryResponse := controller.CategoryService.FindById(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
