package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func GetLoto6Results() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto6Results()
		return c.JSON(http.StatusOK, data)
	}
}
