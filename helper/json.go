package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(http *http.Request, result interface{}) {
	decoder := json.NewDecoder(http.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
