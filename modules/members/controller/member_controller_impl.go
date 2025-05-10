package controller

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/members/domain/service"
	"golang-point-of-sales-system/modules/members/dto/request"
	"golang-point-of-sales-system/pkg/utils/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type MemberControllerImpl struct {
	MemberService service.MemberService
}

func NewMemberController(memberService service.MemberService) MemberController {
	return &MemberControllerImpl{
		MemberService: memberService,
	}
}

func (controller *MemberControllerImpl) Create(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	memberRequest := request.MemberCreateRequest{}
	helper.ReadFromRequestBody(cRequest, &memberRequest)

	memberResponse := controller.MemberService.Create(cRequest.Context(), memberRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) Update(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	memberRequest := request.MemberUpdateRequest{}
	helper.ReadFromRequestBody(cRequest, &memberRequest)
	memberId := params.ByName("memberId")

	parsedUUID, err := uuid.Parse(memberId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	memberRequest.Id = parsedUUID

	memberResponse := controller.MemberService.Update(cRequest.Context(), memberRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) Delete(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	memberId := params.ByName("memberId")

	parsedUUID, err := uuid.Parse(memberId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	controller.MemberService.Delete(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "member deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) FindById(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	memberId := params.ByName("memberId")

	parsedUUID, err := uuid.Parse(memberId)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid UUID format",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	memberResponse := controller.MemberService.FindById(cRequest.Context(), parsedUUID)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MemberControllerImpl) FindAll(writer http.ResponseWriter, cRequest *http.Request, params httprouter.Params) {
	memberResponse := controller.MemberService.FindAll(cRequest.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   memberResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
