package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func GetLoto7Statistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto7Statistics()
		return c.JSON(http.StatusOK, data)
	}
}
