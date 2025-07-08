package Controllers

import (
	"net/http"

	service "loto_depository/src/Services"

	"github.com/labstack/echo/v4"
)

func GetLoto6LatelyStatistics() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := service.GetLoto6LatelyStatistics()
		return c.JSON(http.StatusOK, data)
	}
}
