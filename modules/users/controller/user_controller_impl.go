package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/dto/response"
	"golang-point-of-sales-system/modules/users/domain/service"
	"golang-point-of-sales-system/modules/users/dto/request"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userRequest := request.UserCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &userRequest)

	userResponse := controller.UserService.Create(cRequest.Context(), userRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userRequest := request.UserUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &userRequest)

	userId := params.ByName("userId")

	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	userRequest.Id = parsedUUID

	userResponse := controller.UserService.Update(cRequest.Context(), userRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.UserService.Delete(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	userResponse := controller.UserService.FindById(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userRequest := request.UserLoginRequest{}
	helper.ReadFromRequestBody(cRequest, &userRequest)

	userResponse := controller.UserService.Login(cRequest.Context(), userRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) ChangePassword(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	userRequest := request.UserChangePasswordRequest{}
	helper.ReadFromRequestBody(cRequest, &userRequest)

	userResponse := controller.UserService.ChangePassword(cRequest.Context(), userRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
