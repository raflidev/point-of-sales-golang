package middleware

import (
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/dto/response"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, cRequest *http.Request) {
	if cRequest.Header.Get("X-API-Key") != "RAHASIA" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	} else {
		middleware.Handler.ServeHTTP(writer, cRequest)
	}
}
