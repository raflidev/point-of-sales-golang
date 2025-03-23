package exception

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/dto/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writter, request, err) {
		return
	}

	if validationError(writter, request, err) {
		return
	}

	internalServerError(writter, request, err)
}

func validationError(writter http.ResponseWriter, cRequest *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writter, webResponse)
		return true
	}

	return false
}

func notFoundError(writter http.ResponseWriter, cRequest *http.Request, err interface{}) bool {
	if err == nil {
		return false
	}

	if notFoundError, ok := err.(NotFoundError); ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   notFoundError.Error,
		}

		helper.WriteToResponseBody(writter, webResponse)
		return true
	}

	return false
}

func internalServerError(writter http.ResponseWriter, cRequest *http.Request, err interface{}) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   nil,
	}

	helper.WriteToResponseBody(writter, webResponse)
}
