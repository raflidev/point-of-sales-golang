package adapters

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
)

// HttprouterHandlerToEchoHandler mengkonversi handler dari httprouter ke handler echo
func HttprouterHandlerToEchoHandler(httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params)) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Konversi parameter Echo ke parameter httprouter
		var params httprouter.Params

		// Tambahkan semua parameter dari Echo ke httprouter.Params
		for _, name := range c.ParamNames() {
			params = append(params, httprouter.Param{
				Key:   name,
				Value: c.Param(name),
			})
		}

		// Memanggil handler httprouter
		httpRouterHandler(c.Response().Writer, c.Request(), params)
		return nil
	}
}
